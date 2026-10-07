package methods

// Old names keep working. Registration order does not matter: Alias only records the name,
// and lookup resolves it later; TestAliasesResolve checks every target exists.
func init() {
	Alias("containsString", "contains")
	Alias("replaceAllStringOccurrences", "replaceAll")
	Alias("extractOneByREGEX", "extract")
}
