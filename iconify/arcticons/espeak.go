package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Espeak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.646 9.373C14.756 15.22 9.957 21.266 4.5 25.851c3.93.542 6.57 3.817 8.338 6.981c6.18 11.064 12.812 10.514 19.093.37c2.675-4.321 7.997-9.03 11.57-9.275c-5.29-5.25-10.217-10.794-14.456-16.897c-3.076 6.819-6.192 9.356-9.399 2.343Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.65 23.804c2.589-.553 8.795-8.256 13.822-3.632c1.086 1 3.42 1.313 3.42 1.313c-2.686 4.932-5.425 8.613-8.278 9.51s-5.82-.989-8.964-7.191Z"/>`),
		g.Group(children),
	)
}