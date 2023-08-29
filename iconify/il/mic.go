package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M555 345q-2 50-20 94t-50 78t-72 58t-89 31v51h81q12 0 12 12v69q0 5-3 8t-9 4H150q-5 0-8-4t-3-8v-69q0-12 11-12h81v-51q-48-8-89-31t-72-58t-49-78t-21-94q0-12 12-12h69q10 0 12 11q2 36 17 68t40 55t58 37t70 14t70-14t57-37t41-55t17-68q0-11 11-11h70q5 0 9 3t2 9zm-277 80q-24 0-45-9t-37-25t-25-37t-9-45V124q0-24 9-45t25-37t37-25t45-9t45 9t37 25t25 37t9 45v185q0 24-9 45t-25 37t-37 25t-45 9z"/>`),
		g.Group(children),
	)
}