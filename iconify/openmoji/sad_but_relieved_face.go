package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SadButRelievedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M12.286 36a24 24 0 1 0 24-24a24.027 24.027 0 0 0-24 24Z"/><path fill="#92d3f5" d="M19.683 55.204a6.3 6.3 0 0 1-.495-.02a6.068 6.068 0 0 1-5.56-6.52c.388-4.867 5.223-9.021 5.428-9.196a1.906 1.906 0 0 1 1.42-.46a1.941 1.941 0 0 1 1.331.68c.44.52 4.28 5.194 3.902 9.935a6.02 6.02 0 0 1-2.109 4.133a6.006 6.006 0 0 1-3.917 1.448Z"/><path d="M48.285 35.174a3 3 0 1 1-3-3a3.001 3.001 0 0 1 3 3Zm-18 0a3 3 0 1 1-3-3a3.001 3.001 0 0 1 3 3Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M25.451 56.293a22.963 22.963 0 1 0-11.377-14.299"/><path fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2" d="M22.951 52.995a5.026 5.026 0 0 0 1.762-3.45c.353-4.436-3.503-9.014-3.667-9.207a.952.952 0 0 0-1.343-.107c-.193.165-4.724 4.076-5.078 8.513a5.06 5.06 0 0 0 8.327 4.251Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M42.786 46a12.449 12.449 0 0 0-6.843-1.853A10.389 10.389 0 0 0 29.786 46m21.636-19.702a7.403 7.403 0 0 1-5.304-.324a7.4 7.4 0 0 1-4.11-3.412m-20.823 3.725a8.44 8.44 0 0 0 9.413-3.718"/>`),
		g.Group(children),
	)
}