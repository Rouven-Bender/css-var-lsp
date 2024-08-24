# css-var-lsp

## Limitations
This lsp expects the json rpc messages to send the id as an int while the offical spec would allow for
string or number

Also the lsp only supports lower case letters for Vars currently and no special characters like -
## How to build the project
``go build .
## Connection to neovim
Add this to your init.lua
```lua
local css_var_lsp = vim.lsp.start_client {
    name = "css-var-lsp",
    cmd = { "/path/to/the/binary/for/the/lsp/for/your/os" },
}
if not css_var_lsp then
    vim.notify "hey, the css var lsp didn't launch"
    return
end

vim.api.nvim_create_autocmd("FileType", {
    pattern = "css",
    callback = function()
        vim.lsp.buf_attach_client(0, css_var_lsp)
    end,
})
```
