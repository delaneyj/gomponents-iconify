package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodVessel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 6v19.38l-2.926-2.02l.092-.01l-.966-8.757l-1.987.22l.784 7.112l-6.617-4.568l-1.507 6.027l4.42 3.49l-4.055 1.021l.489 1.94l5.546-1.398l6.35 5.012l.005.005a.95.95 0 0 1 .372.755V42h10V27.113a.94.94 0 0 1 .519-.829l5.151-2.673l7.593 1.847l.473-1.943l-5.438-1.323l5.98-3.248l-2.269-5.293l-7.576 3.597l1.043-4.736l-1.953-.43l-1.38 6.266L28 19.43V6H18Zm5.227 2.85a1 1 0 1 1-1.7 1.055a1 1 0 0 1 1.7-1.055Zm1.678 5.377a1 1 0 1 0-1.055-1.7a1 1 0 0 0 1.055 1.7Zm-2 9a1 1 0 1 0-1.055-1.7a1 1 0 0 0 1.055 1.7Zm1.322 5.622a1 1 0 1 1-1.7 1.056a1 1 0 0 1 1.7-1.056Zm-1.322 10.378a1 1 0 1 0-1.055-1.7a1 1 0 0 0 1.055 1.7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}