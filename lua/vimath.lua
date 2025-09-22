local M = {}

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

	-- Ensure dir
	vim.fn.mkdir(FILES_PATH, "p")
	local file = io.open(FILES_PATH .. FILE_NAME, "w")
	if not file then
		vim.notify("Could not open file: " .. FILES_PATH .. FILE_NAME, vim.log.levels.ERROR)
		return
	end

	for _, line in pairs(lines) do
		file:write(line .. "\n")
	end
	file:close()

	local cmd = "cd " .. PLUGIN_PATH .. " && go run core/*.go --file " .. FILES_PATH .. FILE_NAME
	local output = vim.fn.system(cmd)
	print(output)

	file = io.open(FILES_PATH .. "OUTPUT_" .. FILE_NAME, "r")
	if not file then
		vim.notify("Could not open file: " .. FILES_PATH .. FILE_NAME, vim.log.levels.ERROR)
		return
	end
	file:write(output)
	file:close()
end

return M
