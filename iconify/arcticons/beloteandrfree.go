package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beloteandrfree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m24 36.904l-8.518 3.1a1.893 1.893 0 0 1-2.427-1.131L4.606 15.66a1.893 1.893 0 0 1 1.132-2.427L20.13 7.996a1.893 1.893 0 0 1 2.427 1.131l1.455 3.997"/><path d="m20.401 22.979l-4.24-3.499l-1.098 5.94l3.422 2.824M6.914 15.298l3.56 4.245l-.001-5.54m6.468-2.354l3.56 4.245l-.002-5.54"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="19.102" height="28.488" x="20.642" y="9.756" rx="2" ry="2" transform="rotate(20 30.194 24)"/><path d="M27.661 29.426a.688.688 0 0 0 .915.34l4.92-2.254a3.067 3.067 0 0 0 1.604-1.724a3.057 3.057 0 0 0-.087-2.354a3.052 3.052 0 0 0-1.724-1.602a3.054 3.054 0 0 0-2.314.07a3.064 3.064 0 0 0-1.711-1.57a3.058 3.058 0 0 0-2.35.086a3.057 3.057 0 0 0-1.602 1.726a3.058 3.058 0 0 0 .09 2.353l2.26 4.93ZM25.785 15.27l2.447.89m-.91-5.11l1.467-.227l-1.78 4.893m10.381 3.778l2.447.89m-.91-5.11l1.467-.227l-1.78 4.893"/></g>`),
		g.Group(children),
	)
}