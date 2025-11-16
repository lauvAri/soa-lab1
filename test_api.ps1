# SOA Lab1 API 测试脚本
# 使用方法: PowerShell 中执行 .\test_api.ps1

$baseUrl = "http://localhost:8080/api"

Write-Host "=== SOA Lab1 API 测试 ===" -ForegroundColor Green
Write-Host ""

# 1. 测试获取所有用户（应包含角色信息）
Write-Host "1. 测试 GET /api/users (获取所有用户，包含角色)" -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "$baseUrl/users" -Method GET -ContentType "application/json"
    Write-Host "状态码: $($response.StatusCode)" -ForegroundColor Green
    Write-Host "响应内容:" -ForegroundColor Cyan
    $response.Content | ConvertFrom-Json | ConvertTo-Json -Depth 10
    Write-Host ""
} catch {
    Write-Host "错误: $_" -ForegroundColor Red
    Write-Host ""
}

# 2. 测试根据ID获取用户
Write-Host "2. 测试 GET /api/users/3 (根据ID获取用户，包含角色)" -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "$baseUrl/users/3" -Method GET -ContentType "application/json"
    Write-Host "状态码: $($response.StatusCode)" -ForegroundColor Green
    Write-Host "响应内容:" -ForegroundColor Cyan
    $response.Content | ConvertFrom-Json | ConvertTo-Json -Depth 10
    Write-Host ""
} catch {
    Write-Host "错误: $_" -ForegroundColor Red
    Write-Host ""
}

# 3. 测试创建用户
Write-Host "3. 测试 POST /api/users (创建用户)" -ForegroundColor Yellow
$userData = @{
    name = "测试用户PS"
    roleId = 1
} | ConvertTo-Json

try {
    $response = Invoke-WebRequest -Uri "$baseUrl/users" -Method POST -Body $userData -ContentType "application/json"
    Write-Host "状态码: $($response.StatusCode)" -ForegroundColor Green
    Write-Host "响应内容:" -ForegroundColor Cyan
    $response.Content | ConvertFrom-Json | ConvertTo-Json -Depth 10
    Write-Host ""
} catch {
    Write-Host "错误: $_" -ForegroundColor Red
    Write-Host ""
}

# 4. 测试聚合接口
Write-Host "4. 测试 GET /api/dashboard (聚合接口)" -ForegroundColor Yellow
try {
    $response = Invoke-WebRequest -Uri "$baseUrl/dashboard" -Method GET -ContentType "application/json"
    Write-Host "状态码: $($response.StatusCode)" -ForegroundColor Green
    Write-Host "响应内容:" -ForegroundColor Cyan
    $response.Content | ConvertFrom-Json | ConvertTo-Json -Depth 10
    Write-Host ""
} catch {
    Write-Host "错误: $_" -ForegroundColor Red
    Write-Host ""
}

Write-Host "=== 测试完成 ===" -ForegroundColor Green

