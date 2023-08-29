package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinDrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M277 131q0 27-13 61t-32 63t-37.5 55t-31.5 40l-14 15q-5-5-13.5-15T105 311t-39-56.5t-31-62T21 131q0-53 37.5-90.5T149 3t90.5 37.5T277 131zm-170-.5q0 17.5 12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30zM0 387h299v42H0v-42z"/>`),
		g.Group(children),
	)
}