install_pre_commit:
	mkdir -p .git/hooks
	echo "./gradlew check" > .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit

.PHONY: install_pre_commit
.SILENT: install_pre_commit
