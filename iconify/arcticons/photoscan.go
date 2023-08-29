package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photoscan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31" height="36" x="8.5" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 8.224h24v24H12z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.173 13.197a8.001 8.001 0 1 0 10.935 10.701"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.762 8.222a8 8 0 0 0 11.24 11.24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.866 9.628a8.001 8.001 0 1 0 10.73 10.729"/><circle cx="31.027" cy="13.197" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}