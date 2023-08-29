package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devcontainer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M127.711 477.547c-170.281-98.17-170.281-344.925 0-443.094S512 59.662 512 256S297.993 575.716 127.711 477.547zm123.72-188.214l-89.884-89.884l-28.284 28.284l61.6 61.6l-61.6 61.599l28.284 28.285l89.884-89.884zm131.927-4.889l-61.6-61.6l61.6-61.599l-28.285-28.284l-89.883 89.883l89.883 89.885l28.285-28.285z"/>`),
		g.Group(children),
	)
}