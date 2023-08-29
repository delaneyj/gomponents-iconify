package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 197v124q0 58 36.5 102t93.5 52q69 7 119-37.5T299 327V155q0-22-22-22q-9 0-15 6t-6 16v170q0 47-34.5 79T139 432q-42-4-69-36.5T43 321v-81q0-17-13-30T0 197zM235 69q0-9-6-15t-16-6q-9 0-15 6t-6 15v192h43V69zm-64-42q0-22-22-22q-9 0-15 6t-6 16v234h43V27zm-64 64q0-22-22-22q-9 0-15 6t-6 16v170h43V91z"/>`),
		g.Group(children),
	)
}