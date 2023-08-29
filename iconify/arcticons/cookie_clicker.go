package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookieClicker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.218 43.43A21.415 21.415 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24c0 3.848-1.01 7.46-2.781 10.584"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.878 18.43l-2.242 3.617l1.906 3.532l2.888-.869l3.056-2.916l-1.57-3.729l-4.038.365zm1.683-5.439h4.177l2.047.953l1.402-1.85l-3.505-1.907l-2.635.953l-1.486 1.851zm13.99 8.327L32.617 24l3.336 2.757l1.262-1.654l.505-2.411l1.43-.673l-1.683-2.44l-3 .449l.084 1.29zM14.084 35.056l.925-2.327l1.627.308l.448 1.122l-2.523 1.262l-.477-.365zm6.112 4.57v-3.42s1.823-.785 1.963-.785s1.841.504 1.841.504l.178 1.991l-1.234 2.523l-2.748-.813Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m34.013 41.548l-7.604-7.604a1.397 1.397 0 0 1 0-1.976h0a1.397 1.397 0 0 1 1.976 0l6.636 6.636"/><path d="m35.02 38.604l-2.518-2.518a1.397 1.397 0 0 1 0-1.976h0a1.397 1.397 0 0 1 1.976 0l2.518 2.518"/><path d="m36.996 36.628l-1.82-1.82a1.397 1.397 0 0 1 0-1.976h0a1.397 1.397 0 0 1 1.975 0l1.821 1.82"/><path d="m38.972 34.653l-1.059-1.06a1.397 1.397 0 0 1 0-1.975h0a1.397 1.397 0 0 1 1.976 0l2.027 2.027c1.337 1.337 2.112 3.235 3.584 4.707L38.352 45.5c-1.55-1.705-5.009-2.062-6.934-2.324c-1.092-.149-1.376-.678-1.22-1.686s1.704-1.162 2.982-.774"/></g>`),
		g.Group(children),
	)
}