package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fuelphp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 1024q-32 0-64-6q-32 6-64 6q-87 0-160.5-43T43 864.5T0 704q0-72 34.5-173T111 347t97.5-174t77-123T320 0q27 38 64 96l64-96q13 18 34.5 50T559 173t97.5 174t77 184T768 704q0 87-43 160.5T608.5 981T448 1024zM320 96q-10 16-27.5 43.5t-61.5 105t-78 147t-61.5 152T64 680q0 116 75 198t181 82t181-82t75-198q0-54-26.5-134.5t-64-156t-75.5-143t-63.5-109T320 96zm96 51q224 365 224 557q0 71-29 133.5T530 945q77-28 125.5-101T704 680q0-54-26.5-134.5t-64-156t-75-143T474 137l-26-41q-13 20-32 51z"/>`),
		g.Group(children),
	)
}