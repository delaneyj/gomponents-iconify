package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoNpm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M227.6 213.1H256v57.1h-28.4z"/><path fill="currentColor" d="M0 156v171.4h142.2V356H256v-28.6h256V156Zm142.2 142.9h-28.4v-85.7H85.3v85.7H28.4V184.6h113.8Zm142.2 0h-56.9v28.6h-56.9V184.6h113.8Zm199.2 0h-28.4v-85.7h-28.4v85.7h-28.4v-85.7H370v85.7h-56.9V184.6h170.7v114.3Z"/>`),
		g.Group(children),
	)
}