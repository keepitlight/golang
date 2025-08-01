package hotfix

// History is the history provider for hotfix.
//
// History 是补丁程序执行的历史记录提供者接口
type History interface {
	// Fixed returns the last fixed version.
	//
	// Fixed 返回最近一次修复的版本号
	Fixed() Version
	// Record records the fixed version.
	//
	// Record 记录修复的版本号
	Record(v Version, summary []string)
}
