package purple

import (
	"github.com/avence12/licenses/testdata/src/colors/broken"
	"github.com/avence12/licenses/testdata/src/colors/red"
)

func purple() string {
	return "purple" + broken.broken() + red.red()
}
