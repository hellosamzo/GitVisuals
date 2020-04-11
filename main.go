package main

import (
	"flag"
)

// scan a path and its children directories
// search for git repos
func scan(dir string) {
	fmt.Printf("Found dirs:\n\n")
	repos = := recursiveScanFolder(dir)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repos)
	fmt.Printf("n\nSuccessfully added\n\n")
}

// stats of commits etc
func stats(email string) {
	print("stats")
}

func recursiveScanFolder(dir string) []string {
	return scanGitDirectories(make([]string, 0), dir)
}

// scan for git directories
func scanGitDirectories(directories []string, dir string) []string {
	dir = strings.TrimSuffix(dir, "/")

	f, err := os.Open(dir)

	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
        if file.IsDir() {
            path = dir + "/" + file.Name()
            if file.Name() == ".git" {
                path = strings.TrimSuffix(path, "/.git")
                fmt.Println(path)
                folders = append(directories, path)
                continue
            }
            if file.Name() == "vendor" || file.Name() == "node_modules" {
                continue
            }
            folders = scanGitFolders(directories, path)
        }
	}
	
	return directories

}

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
}

func main() {
	var dir string
	var email string
	flag.StringVar(&dir, "add", "", "add a new directory to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the emai to scan")
	flag.Parse()

	if dir != "" {
		scan(dir)
		return
	}

	stats(email)
}
