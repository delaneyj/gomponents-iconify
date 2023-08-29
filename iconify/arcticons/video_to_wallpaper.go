package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoToWallpaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.29 22.24h-5.282m-7.854 0h-.677v14.895h24.917V22.24h-2.031m6.77 11.781l8.667 5.281V19.938l-8.667 5.281h-4.739v8.802h4.739z"/><circle cx="23.894" cy="16.146" r="7.448" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.946" cy="18.719" r="5.146" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23.894" cy="16.146" r="2.844" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.946" cy="18.719" r="1.76" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}