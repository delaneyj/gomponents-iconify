package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.407 4.5h9.185l6.664 7.094l.621 11.236l-4.916 7.778l-1.351 3.925v6.824L24 43.5l-5.61-2.143v-6.824l-1.35-3.925l-4.917-7.778l.62-11.235Zm0 0v8.255m-6.663-1.161l6.663 1.162M12.122 22.83l7.285-10.074M28.593 4.5l-9.185 8.255M28.593 4.5l2.764 13.36m3.899-6.266l-3.899 6.266m4.521 4.97l-4.52-4.97m-.398 12.748l.397-12.748m-11.95-5.104l11.95 5.104"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.122 22.83l10.842 1.819l7.997 5.96m.396-12.749l-8.393 6.788m-3.557-11.892l3.557 11.892m-5.924 5.96l5.925-5.96"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.964 24.648l-4.574 9.885l12.571-3.925M18.39 34.533h11.22"/>`),
		g.Group(children),
	)
}