package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KvaesitsoIcons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.983 20.862h5.638V26.5h-5.638zm-4.983 0l2.685 5.638h-5.638L24 20.862zM8.636 15.759l13.46-6.715a4.152 4.152 0 0 1 3.66-.023L39.365 15.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.926 19.778v7.807c0 .958.777 1.734 1.735 1.734h24.287c.958 0 1.734-.776 1.734-1.734v-7.807c0-.958-.776-1.735-1.735-1.735H11.661c-.958 0-1.735.777-1.735 1.735Zm25.572 22.388v-7.863a2 2 0 0 0-2-2H14.11a2 2 0 0 0-2 2v7.61"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.929" cy="23.681" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}