package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.23 2.503a9.5 9.5 0 0 0-.46 18.994c.338.008.643-.086.898-.28c.248-.187.42-.444.539-.72c.235-.54.314-1.249.318-1.98c.015-2.641 2.332-4.935 4.948-4.836c.72.027 1.413-.012 1.948-.188c.27-.09.534-.225.736-.436c.212-.221.332-.503.34-.827a9.5 9.5 0 0 0-9.267-9.727Zm-3.34 12.9a1 1 0 1 0 .243 1.985a1 1 0 0 0-.243-1.985Zm-3.342-2.614a1 1 0 1 1 1.985-.243a1 1 0 0 1-1.985.243ZM7.92 7.462a1 1 0 1 0 .243 1.985a1 1 0 0 0-.243-1.985Zm7.778 1.819a1 1 0 1 1 1.986-.243a1 1 0 0 1-1.986.243Zm-3.704 1.598a1 1 0 1 0 .242 1.985a1 1 0 0 0-.242-1.985Zm-.6-3.842a1 1 0 1 1 1.985-.243a1 1 0 0 1-1.985.243Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}