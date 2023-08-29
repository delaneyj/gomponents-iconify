package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.463 20.55l-4.387 4.118v1.271l-2.338 2.469h1.302L20.957 42.31a66.098 66.098 0 0 0-6.755 1.189m13.261-14.075l-1.526 1.527l1.526 1.527"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.077 25.94l-6.722-6.723l2.944-2.944l6.081 6.08m-3.527 7.722l-5.644 5.645l3.069 3.07l2.062-2.063m-2.776-2.997l2.695 2.695M18.69 21.49l2.758-2.758m-.598 4.894l2.84-2.84m3.773-.236l4.388 4.118v1.271l2.338 2.469h-1.302L33.97 42.31s3.603.43 6.545 1.138M27.463 29.424l1.527 1.527l-1.527 1.527"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.85 25.94l6.722-6.723l-2.944-2.944l-6.08 6.08m3.526 7.722l5.644 5.645l-3.069 3.07l-2.062-2.063m2.776-2.997l-2.695 2.695m2.568-14.935l-2.757-2.758m.598 4.894l-2.841-2.84m3.068-9.73a3.652 3.652 0 0 0-1.292-1.824l-4.802-3.94a3.14 3.14 0 0 0-1.966-.786M7.485 20.724s4.284.64 7.485-4.22s10.581-3.267 13.24-2.322"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.648 13.244a4.26 4.26 0 0 0-.44 1.763c.102.616-.937 1.253-.937 1.253s1.584.12 2.096.242c.987.235 2.403.055 1.966-2.867"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.333 13.635s.21-2.089 1.39-1.23s1.195-1.6 2.215-2.792c.534-.33 1.267-.141 1.91-.186M29.107 6.13c.373 1.347.056 2.435-.472 1.836a1.283 1.283 0 0 0-1.548-.56a4.106 4.106 0 0 0-2.381 3.221a12.646 12.646 0 0 1-.058 2.617"/>`),
		g.Group(children),
	)
}