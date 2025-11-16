# 用户接口完整测试脚本
Write-Host "=== 用户接口完整测试 ===" -ForegroundColor Green
Write-Host ""

$baseUrl = "http://localhost:8080/api/users"

# 1. 获取所有用户
Write-Host "1. GET /api/users (获取所有用户)" -ForegroundColor Yellow
try {
    $r = Invoke-WebRequest -Uri $baseUrl -Method GET -UseBasicParsing
    Write-Host "   Status: $($r.StatusCode) - OK" -ForegroundColor Green
    $json = $r.Content | ConvertFrom-Json
    Write-Host "   返回用户数: $($json.data.Count)" -ForegroundColor Cyan
    if ($json.data.Count -gt 0) {
        Write-Host "   第一个用户: ID=$($json.data[0].id), 名称=$($json.data[0].name), 角色=$($json.data[0].role.roleName)" -ForegroundColor Cyan
    }
} catch {
    Write-Host "   Status: ERROR - $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 2. 根据ID获取用户
Write-Host "2. GET /api/users/3 (根据ID获取用户)" -ForegroundColor Yellow
try {
    $r = Invoke-WebRequest -Uri "$baseUrl/3" -Method GET -UseBasicParsing
    Write-Host "   Status: $($r.StatusCode) - OK" -ForegroundColor Green
    $json = $r.Content | ConvertFrom-Json
    Write-Host "   用户ID: $($json.data.id), 名称: $($json.data.name), 角色: $($json.data.role.roleName)" -ForegroundColor Cyan
    Write-Host "   role字段存在: $($json.data.role -ne $null)" -ForegroundColor Cyan
} catch {
    Write-Host "   Status: ERROR - $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 3. 根据角色ID获取用户
Write-Host "3. GET /api/users/role/1 (根据角色ID获取用户)" -ForegroundColor Yellow
try {
    $r = Invoke-WebRequest -Uri "$baseUrl/role/1" -Method GET -UseBasicParsing
    Write-Host "   Status: $($r.StatusCode) - OK" -ForegroundColor Green
    $json = $r.Content | ConvertFrom-Json
    Write-Host "   返回用户数: $($json.data.Count)" -ForegroundColor Cyan
    if ($json.data.Count -gt 0) {
        Write-Host "   role字段存在: $($json.data[0].role -ne $null)" -ForegroundColor Cyan
        if ($json.data[0].role) {
            Write-Host "   角色信息: ID=$($json.data[0].role.id), 名称=$($json.data[0].role.roleName)" -ForegroundColor Cyan
        }
    }
} catch {
    Write-Host "   Status: ERROR - $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 4. 创建用户
Write-Host "4. POST /api/users (创建用户)" -ForegroundColor Yellow
$timestamp = Get-Date -Format 'HHmmss'
$userData = @{
    name = "测试用户$timestamp"
    roleId = 1
} | ConvertTo-Json -Compress

$newUserId = $null
try {
    $r = Invoke-WebRequest -Uri $baseUrl -Method POST -Body $userData -ContentType "application/json" -UseBasicParsing
    Write-Host "   Status: $($r.StatusCode) - Created" -ForegroundColor Green
    $json = $r.Content | ConvertFrom-Json
    $newUserId = $json.data.id
    Write-Host "   新用户ID: $newUserId, 名称: $($json.data.name), 角色: $($json.data.role.roleName)" -ForegroundColor Cyan
} catch {
    Write-Host "   Status: ERROR - $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

# 5. 更新用户
if ($newUserId) {
    Write-Host "5. PUT /api/users/$newUserId (更新用户)" -ForegroundColor Yellow
    $updateData = @{
        name = "更新用户$timestamp"
        roleId = 2
    } | ConvertTo-Json -Compress
    
    try {
        $r = Invoke-WebRequest -Uri "$baseUrl/$newUserId" -Method PUT -Body $updateData -ContentType "application/json" -UseBasicParsing
        Write-Host "   Status: $($r.StatusCode) - OK" -ForegroundColor Green
        $json = $r.Content | ConvertFrom-Json
        Write-Host "   更新后名称: $($json.data.name), 角色: $($json.data.role.roleName)" -ForegroundColor Cyan
    } catch {
        Write-Host "   Status: ERROR - $($_.Exception.Message)" -ForegroundColor Red
    }
    Write-Host ""
} else {
    Write-Host "5. PUT /api/users/{id} (跳过 - 未创建用户)" -ForegroundColor Gray
    Write-Host ""
}

# 6. 删除用户
if ($newUserId) {
    Write-Host "6. DELETE /api/users/$newUserId (删除用户)" -ForegroundColor Yellow
    try {
        $r = Invoke-WebRequest -Uri "$baseUrl/$newUserId" -Method DELETE -UseBasicParsing
        Write-Host "   Status: $($r.StatusCode) - OK" -ForegroundColor Green
        $json = $r.Content | ConvertFrom-Json
        Write-Host "   删除结果: $($json.message)" -ForegroundColor Cyan
    } catch {
        Write-Host "   Status: ERROR - $($_.Exception.Message)" -ForegroundColor Red
    }
    Write-Host ""
} else {
    Write-Host "6. DELETE /api/users/{id} (跳过 - 未创建用户)" -ForegroundColor Gray
    Write-Host ""
}

# 7. 测试错误路径
Write-Host "7. 测试错误路径 /api/users/users (应该返回错误)" -ForegroundColor Yellow
try {
    $r = Invoke-WebRequest -Uri "$baseUrl/users" -Method GET -UseBasicParsing
    Write-Host "   Status: $($r.StatusCode) (意外成功)" -ForegroundColor Yellow
} catch {
    $status = if ($_.Exception.Response) { $_.Exception.Response.StatusCode.value__ } else { "Error" }
    $color = if ($status -eq 500) { "Yellow" } elseif ($status -eq 404) { "Green" } else { "Red" }
    Write-Host "   Status: $status (预期错误)" -ForegroundColor $color
    if ($status -eq 500) {
        Write-Host "   错误类型: NumberFormatException (路径匹配错误)" -ForegroundColor Gray
    }
}
Write-Host ""

Write-Host "=== 测试完成 ===" -ForegroundColor Green

