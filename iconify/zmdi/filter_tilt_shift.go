package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterTiltShift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M191 47q-46 5-83 34L77 51q50-41 114-47v43zm156 4l-30 30q-38-29-83-34V4q63 6 113 47zm34 144q-5-46-34-84l30-30q41 50 48 114h-44zM78 111q-29 38-35 84H0q6-64 47-114zM43 237q6 46 35 83l-31 31Q6 301 0 237h43zm233-21q0 27-18.5 45.5t-45 18.5t-45.5-18.5t-19-45.5t19-45.5t45.5-18.5t45 18.5T276 216zm71 105q29-38 34-83h44q-7 63-48 113zm-113 64q46-6 83-34l30 30q-50 41-113 47v-43zm-157-4l31-30q37 29 83 34v43q-64-6-114-47z"/>`),
		g.Group(children),
	)
}