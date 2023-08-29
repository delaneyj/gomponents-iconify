package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vault(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#ffd814" d="m0 1.953l63.76 124.094L128 1.953Zm53.841 49.254H43.684V41.06H53.84zm0-15.227H43.684V25.822H53.84ZM69.08 66.444H58.97V56.286h10.108zm0-15.237H58.97V41.06h10.108zm0-15.227H58.97V25.822h10.108Zm15.147 15.227h-10.2V41.06h10.159Zm-10.2-15.227V25.822h10.159V35.98z"/>`),
		g.Group(children),
	)
}