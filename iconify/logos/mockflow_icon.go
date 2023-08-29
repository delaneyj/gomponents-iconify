package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MockflowIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#7957FF" d="m146.286 152.381l73.143-51.403h-73.143z"/><path fill="#1B7CE2" d="M90.006 84.114L51.81 141.206l-3.048 4.877h-.203L0 219.429V0z"/><path fill="#7957FF" d="M219.429 48.762h-73.143V0H256z"/><path fill="#F845C6" d="M146.286 0v219.429l-97.524-73.346l3.048-4.877l38.196-57.092z"/>`),
		g.Group(children),
	)
}