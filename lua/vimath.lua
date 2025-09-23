local M = {}
local ns_id = vim.api.nvim_create_namespace("demo")

function M.setup()
	vim.api.nvim_create_user_command("VimathRun", function()
		M.get_current_buffer_content()
	end, {})
end

function M.get_current_buffer_content()
	local bufnr = vim.api.nvim_get_current_buf()
	local line_count = vim.api.nvim_buf_line_count(bufnr)
	local lines = vim.api.nvim_buf_get_lines(bufnr, 0, line_count, false)

	-- TODO: implement better file/plugin DIRs strucutre
	-- consider not using a temp file and sending buffer
	-- directly to the interpreter binary
	local PLUGIN_PATH = vim.fn.expand("LOCAL_DEV_DIR_DO_NOT_COMMIT")
	local FILES_PATH = PLUGIN_PATH .. "tmp/"
	local FILE_NAME = "test.txt"
	local CURR_FILE_NAME = FILES_PATH .. FILE_NAME

	-- Ensure dir
	vim.fn.mkdir(FILES_PATH, "p")
	local file = io.open(CURR_FILE_NAME, "w")
	if not file then
		vim.notify("Could not open file: " .. CURR_FILE_NAME, vim.log.levels.ERROR)
		return
	end

	for _, line in pairs(lines) do
		file:write(line .. "\n")
	end
	file:close()

	local cmd = "cd " .. PLUGIN_PATH .. " && go run core/*.go --file tmp/" .. FILE_NAME
	local output = vim.fn.system(cmd)
	M.render_results(output, bufnr)

	CURR_FILE_NAME = FILES_PATH .. "OUTPUT_" .. FILE_NAME
	file = io.open(CURR_FILE_NAME, "w")
	if not file then
		vim.notify("Could not open file: " .. CURR_FILE_NAME, vim.log.levels.ERROR)
		return
	end
	file:write(output)
	file:close()
end

function M.render_results(output, bufnr)
	for line in output:gmatch("[^\r\n]+") do
		if not line:find("^%[") then
			local index, val = line:match("(%S+)%s+(%S+)")
			local i = tonumber(index)
			local v = tonumber(val)

			if i and v then
				vim.api.nvim_buf_set_extmark(bufnr, ns_id, i, 0, {
					virt_text = { { tostring(v), "TODO" } },
					virt_text_pos = "right_align",
				})
			end
		end
	end
end

return M
