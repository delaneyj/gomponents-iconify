package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5C9.571 5 8 6.54 8 8c0 .804.28 1.362.865 1.818c.631.492 1.665.897 3.224 1.182H19a1 1 0 1 1 0 2h-2.189c.788.794 1.189 1.803 1.189 3c0 2.958-2.906 5-6 5c-1.998 0-3.827-.814-4.936-2.149a1 1 0 1 1 1.538-1.278C9.285 18.393 10.52 19 12 19c2.429 0 4-1.54 4-3c0-.804-.28-1.362-.865-1.818c-.631-.492-1.665-.897-3.224-1.182H5a1 1 0 1 1 0-2h2.189C6.401 10.206 6 9.197 6 8c0-2.958 2.906-5 6-5c1.477 0 2.852.444 3.915 1.205a1 1 0 0 1-1.164 1.627C14.046 5.327 13.084 5 12 5z"/>`),
		g.Group(children),
	)
}