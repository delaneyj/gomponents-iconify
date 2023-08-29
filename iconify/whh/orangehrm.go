package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orangehrm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M544 256h-32q0-97 135-167q32-17 57-25q-45 0-92 19.5t-82 48t-57.5 61.5t-24.5 60q-33-5-48.5-18.5T384 192q0-35 45.5-80.5T544 33T672 0q55 0 95 8t53 16l12 8q0 12-32 48.5T724.5 156t-95 69.5T544 256zm-32 256q-14-7-39-17.5T380.5 466T256 448q-52 0-90 50.5T128 640q0 62 54.5 122T319 859t161 37q77 0 144.5-37.5T730 762t38-122q0-173-186-270q-57-30-134-50q126 0 226 35t161 108.5T896 640q0 104-60 192.5t-163 140t-225 51.5t-225-51.5t-163-140T0 640q0-69 21.5-120.5t59-80t81-42T256 384q63 0 122.5 18.5T475 451t37 61z"/>`),
		g.Group(children),
	)
}