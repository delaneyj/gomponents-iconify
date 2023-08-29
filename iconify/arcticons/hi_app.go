package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HiApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.026 6.049c-5.698.917-10.01 5.602-10.01 11.277c0 6.338 5.378 11.44 12.058 11.44l3.854-.001c6.68 0 12.056-5.103 12.056-11.44c0-3.877-2.013-7.292-5.092-9.388m-1.74-.99a12.723 12.723 0 0 0-3.656-1.073c.365-2.787-9.44-4.145-6.652 1.893c1.403 3.037 10.555 4.852 13.244-2.327c-4.059 3.538-8.648 3.114-9.81.832"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.294 20.972c-.668 4.48-6.724 3.385-7.376.417c3.144.503 4.924.434 7.376-.417Zm4.125-2.36c0 .903-.571 1.635-1.276 1.635h0c-.704 0-1.275-.732-1.275-1.635h0c0-.903.57-1.635 1.275-1.635s1.276.732 1.276 1.635c0 0 0 0 0 0Zm-12.651.107c0 .903-.57 1.635-1.275 1.635h0c-.705 0-1.276-.732-1.276-1.635h0c0-.903.571-1.635 1.276-1.635c.704 0 1.275.732 1.275 1.635c0 0 0 0 0 0Zm9.235 9.735c2.738 1.83 2.56 5.458 2.683 8.752c-2.933 2.649-7.377 2.648-11.593 1.008l1.02-2.72"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.244 31.624c2.965-1.245 6.181.45 4.214 2.52c-1.565 1.648-5.204 1.668-7.012.66c-2.03-1.13.733-5.376 2.798-3.18Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.77 28.622c-1.096 1.344-1.09 3.88-1.004 4.732m12.194-2.87c1.775-.803 3.602-1.22 5.325-2.418c.434-.302 1.342-1.418.745-2.643c-.426-.876-1.065-.75-1.703-.693"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.691 26.865c.262.419.798 1.392.35 2.415m-12.443 9.197c.245.182-.99 4.275-.87 5.05c1.564 1.07 3.424 1.1 5.362.856c.374-1.328.801-2.805 1.873-3.872"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.395 41.465l.085 2.123c.68.936 3.392.617 4.775-.55c-.168-1.591-.272-3.268-.218-5.09"/>`),
		g.Group(children),
	)
}