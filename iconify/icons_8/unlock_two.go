package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3c-3.036 0-5.584 1.966-6.625 4.625l1.844.75C11.977 6.435 13.835 5 16 5c2.755 0 5 2.245 5 5v1.375C19.525 10.515 17.826 10 16 10c-5.51 0-10 4.49-10 10s4.49 10 10 10s10-4.49 10-10a9.968 9.968 0 0 0-3-7.125V10c0-3.845-3.155-7-7-7zm0 9c4.43 0 8 3.57 8 8s-3.57 8-8 8s-8-3.57-8-8s3.57-8 8-8zm0 6a2 2 0 0 0-2 2c0 .74.403 1.373 1 1.72V25h2v-3.28c.597-.347 1-.98 1-1.72a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}