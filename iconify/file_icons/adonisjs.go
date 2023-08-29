package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adonisjs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 256c0 206.466 49.534 256 256 256s256-49.534 256-256S462.466 0 256 0S0 49.534 0 256zm173.626 160.582c-58.99.278-91.98-54.306-70.393-107.414l80.305-182.475c19.73-56.102 111.542-71.358 144.938 0l80.29 182.475c21.716 52.478-11.158 108.055-70.392 107.414c-19.592 1.368-58.414-17.567-82.434-16.068c-32.51-.363-54.67 16.891-82.314 16.068zm160.582-78.725L256 158.499l-79.257 179.358c47.532-20.986 108-22.252 157.465 0z"/>`),
		g.Group(children),
	)
}