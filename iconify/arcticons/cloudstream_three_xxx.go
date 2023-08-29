package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudstreamThreeXxx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.759 29.064v7.086l5.037-3.543Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.294 34.864h-1.07a2.412 2.412 0 0 0-2.52-3.22H35.29s.065-5.82-6.856-5.82c-1.077-1.818-3.129-3.005-6.114-3.005c-5.627 0-6.611 5.59-6.611 5.59a5.682 5.682 0 0 0-7.003 7.842H6.904a1.915 1.915 0 0 0 0 3.83c.015 0 .029-.004.044-.005l31.561.081a1.08 1.08 0 0 0 1.078-1.084v-.007l-.003-1.186h.71a1.509 1.509 0 0 0 .002-3.017h-.002Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.403 32.143c1.103-.333 1.523-2.014 1.523-2.014a7.751 7.751 0 0 1-2.991.445s5.76-1.632 1.211-8.579c.074 3.51-3.27 4.42-6.431 4.398m-3.774-1.273c-.4-1.443 1.33-1.913 1.33-4.83s-4.153-4.227-4.153-4.227s.92 2.62-3.36 2.62s-4.947-1.36-4.947-2.3c0-2.052 3.51-2.793 1.261-5.735c-.37 2.077-2.843 2.942-4.672 3.412s-3.51 1.468-3.51 3.824s2.447 3.543 2.447 5.397s-2.67 2.522-2.67 2.522c.47-1.137-1.557-2.04-1.557-2.04c0 2.386-3.61 2.065-3.61 4.735s3.731 4.231 3.731 4.231"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.551 23.825c0-1.446-1.483-1.595-1.483-2.98M16.453 7.842a1.672 1.672 0 0 0 .309 1.249m1.978.408a3.422 3.422 0 0 1 1.558 2.41M43.251 26.99c.742 1.557-.148 4.171-3.04 5.358m-5.433-13.665a1.974 1.974 0 0 1 1.52 1.186m-1.641 4.543a2.273 2.273 0 0 0 1.762-1.798m-12.525-9.316a2.291 2.291 0 0 0 1.399 1.323m3.509 1.713c.667.31 3.176 2.064 2.36 4.734"/>`),
		g.Group(children),
	)
}