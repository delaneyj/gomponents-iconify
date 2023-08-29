package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoneBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.29 5.79c-.28-1.014.104-2.309.847-3.052A2.519 2.519 0 0 1 17.7 6.3a2.519 2.519 0 0 1 3.562 3.563c-.743.743-2.038 1.128-3.052.848c-.536-.149-1.185-.173-1.579.22L13.07 7.37c.393-.394.369-1.043.22-1.58Zm-5.921 7.279c-.394.393-1.043.369-1.58.22c-1.013-.28-2.308.105-3.051.848A2.519 2.519 0 0 0 6.3 17.7a2.519 2.519 0 0 0 3.563 3.562c.743-.743 1.128-2.038.848-3.052c-.149-.536-.173-1.185.22-1.579L7.37 13.07Z" clip-rule="evenodd"/><path d="m10.931 16.631l5.7-5.7L13.07 7.37l-5.7 5.7l3.563 3.562Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}