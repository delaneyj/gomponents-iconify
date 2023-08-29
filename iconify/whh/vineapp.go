package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vineapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 512v125q-35 3-64 3h-32q-35 74-81.5 140.5t-89 109T547 964t-65.5 46t-33.5 14q-9 0-36-15.5t-68.5-49t-87-81t-91.5-118T82.5 609T23 419T0 192h128q0 92 17 176.5t43.5 147T248 632t66 89.5t62.5 61.5t48 37.5T448 832q12 0 43.5-22t64.5-53.5t58.5-73T640 608q-4-2-11-6t-27-19.5t-38-34.5t-40-52t-38-70t-27-90t-11-112q0-62 36.5-115T572 28T672 0q103 0 163.5 55.5T896 192q0 192-32 256H736q32-128 32-224q0-41-28-68.5T672 128q-41 0-68.5 26T576 224q0 38 8.5 75t31 76.5t57 69t92 48.5T896 512z"/>`),
		g.Group(children),
	)
}