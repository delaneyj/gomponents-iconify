package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiMycenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.822 20.88l-.054 4.772"/><rect width="23.318" height="10.828" x="12.421" y="17.568" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.028"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 24a18.5 18.5 0 1 1-37 0c0-2.554 1.276-8.212 3.17-9.935c0 0-.59-6.827.426-7.596c.862-.652 6.899 1.235 6.899 1.235A14.379 14.379 0 0 1 24 5.5a14.417 14.417 0 0 1 8.404 2.13s5.795-2.1 6.934-1.094s.37 8.028.37 8.028A15.494 15.494 0 0 1 42.5 24Zm-12.483-3.12l-.054 4.772"/>`),
		g.Group(children),
	)
}