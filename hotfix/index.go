package hotfix

import (
	"context"
)

var (
	hotfixes = &fixes{}
)

// SetHistoryProvider to set the history provider.
//
// 设置历史记录接口
func SetHistoryProvider(history History) {
	hotfixes.history = history
}

// Patch to run hotfixes if and only if Do is being called for the first time.
// Argument version indicates the version from which to start patching.
//
// 如果是第一次调用 Do，则运行热修补。 参数 version 指示从指定的版本开始执行修补。
func Patch(ctx context.Context, version Version) error {
	return hotfixes.do(ctx, version)
}

// FixIt to register a hotfix function, which will be executed when the version is reached.
// version must be greater than 0
//
// 注册热修补函数
func FixIt(version Version, title string, hotfix Hotfix) {
	hotfixes.append(version, title, hotfix)
}
