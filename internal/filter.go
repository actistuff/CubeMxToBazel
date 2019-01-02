package cubemxtobazelinternal

// FileFilter can be used to recursively filter a slice of files
type FileFilter interface {
	AssemblyFiles() FileFilter
	SourceFiles() FileFilter
	HeaderFiles() FileFilter
	Condition(string) FileFilter
	InvCondition(string) FileFilter
	Files() MxFiles
}

// MxFiles slice of CubeMx project and driver files
type MxFiles []MxFile

// AssemblyFiles returns all assembly files generated by CubeMx in the file list
func (files MxFiles) AssemblyFiles() FileFilter {
	var filtered MxFiles
	for _, file := range files {
		if file.Category == "sourceAsm" {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

// SourceFiles returns all source files generated by CubeMx in the file list
func (files MxFiles) SourceFiles() FileFilter {
	var filtered MxFiles
	for _, file := range files {
		if file.Category == "source" || file.Category == "sourceC" {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

// HeaderFiles files returns all header file generated by CubeMx in the file list
func (files MxFiles) HeaderFiles() FileFilter {
	var filtered MxFiles
	for _, file := range files {
		if file.Category == "header" {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

// Condition returns all files generated by CubeMx that meet the conditional element
func (files MxFiles) Condition(cond string) FileFilter {
	var filtered MxFiles
	for _, file := range files {
		if file.Condition == cond {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

// InvCondition returns all files generated by CubeMx that do not meet the conditional element
func (files MxFiles) InvCondition(cond string) FileFilter {
	var filtered MxFiles
	for _, file := range files {
		if file.Condition != cond && file.Condition != "" {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

// Files returns the resulting filtered files
func (files MxFiles) Files() MxFiles {
	return files
}
