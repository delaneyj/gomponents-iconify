package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bugjaeger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 1 2.5 24A21.5 21.5 0 0 1 24 2.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5a23.447 23.447 0 0 1 8.281 5.576a16.085 16.085 0 0 0-7.57 4.946a13.76 13.76 0 0 1 11.373.321c.18.34.354.683.517 1.033a13.657 13.657 0 0 0-8.328 4.633a11.67 11.67 0 0 1 10.02.451a23.435 23.435 0 0 1 .51 4.806c0 .704-.036 1.4-.097 2.088a21.421 21.421 0 0 0-29.411 0a23.7 23.7 0 0 1-.098-2.088a23.435 23.435 0 0 1 .51-4.806a11.67 11.67 0 0 1 10.02-.451a13.657 13.657 0 0 0-8.327-4.633c.163-.35.337-.693.517-1.033a13.76 13.76 0 0 1 11.373-.32a16.085 16.085 0 0 0-7.57-4.947A23.447 23.447 0 0 1 24 2.5m-2.734 42.828v-2.296m-5.212.951v-2.068"/><circle cx="16.798" cy="31.862" r="3.478" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.798 28.383v-1.521m2.459 2.54l1.076-1.076m-.057 3.536h1.522m-2.541 2.459l1.076 1.076m-3.535-.057v1.522m-2.46-2.541l-1.076 1.076m.058-3.535h-1.522m2.54-2.46l-1.076-1.076m13.472 17.002v-2.296m5.212.951v-2.068"/><circle cx="31.202" cy="31.862" r="3.478" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.202 28.383v-1.521m-2.459 2.54l-1.076-1.076m.057 3.536h-1.522m2.541 2.459l-1.076 1.076m3.535-.057v1.522m2.46-2.541l1.076 1.076m-.058-3.535h1.522m-2.54-2.46l1.076-1.076"/>`),
		g.Group(children),
	)
}