local-module:
	echo "replace github.com/openopsdev/go-cli-commons => ../go-cli-commons" >> go.mod

remote-module:
	sed -i -e "/replace\ github\.com\/openopsdev\/go-cli-commons\ =>/d" go.mod