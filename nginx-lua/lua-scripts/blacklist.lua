local cjson = require "cjson"

M = {}
local blacklist_dictionary = {}
blacklist_dictionary["172.256.0.1"] = 1
blacklist_dictionary["39.144.1"] = 1
blacklist_dictionary["Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405"] = 1

function M.interrupt()
    ngx.req.read_body()

    -- ngx.log(ngx.DEBUG,"Enter Interupt")

    local data_body = ngx.req.get_body_data()
    if not data_body then
        ngx.log(ngx.DEBUG,"Nil Body")
        return ngx.exit(403)
    end

    local body_value = cjson.decode(data_body)

    local is_match = M.checkBlacklistIP(body_value["ip_address"])
    -- ngx.log(ngx.DEBUG,"is_match 1 = ", is_match)
    if is_match == true then
        ngx.status = ngx.HTTP_FORBIDDEN
        return ngx.exit(ngx.HTTP_FORBIDDEN)
    end

    local is_match = M.checkBlacklistUserAgent(body_value["user_agent"])
    -- ngx.log(ngx.DEBUG,"is_match 2 = ", is_match)
    if is_match == true then
        ngx.status = ngx.HTTP_FORBIDDEN
        return ngx.exit(ngx.HTTP_FORBIDDEN)
    end

    ngx.status = ngx.HTTP_OK
    return ngx.exit(ngx.HTTP_OK)

end

function M.checkBlacklistIP(ip)
    if blacklist_dictionary[ip] == 1  then
        return true
    else
        return false
    end
end

function M.checkBlacklistUserAgent(ua)
    if blacklist_dictionary[ua] == 1  then
        return true
    else
        return false
    end
end


return M