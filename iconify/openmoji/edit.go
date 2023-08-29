package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M12.854 16.577h42.498v42.246H12.854z"/><path fill="#F4AA41" d="m32.775 39.406l4.719-1.782l-2.937-2.937z"/><path fill="#A57939" d="m56.611 11.678l3.784 3.785l-22.014 22.013l-3.784-3.784z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M54.557 28.084v30.684c0 .55-.45 1-1 1H13.168c-.55 0-1-.45-1-1V18.379a1 1 0 0 1 1-1h31.023"/><path d="m38.053 37.749l21.46-21.461a1 1 0 0 0 0-1.415l-2.335-2.335a1 1 0 0 0-1.414 0l-21.462 21.46l-2.625 6.387l6.376-2.636l-3.75-3.75"/></g>`),
		g.Group(children),
	)
}