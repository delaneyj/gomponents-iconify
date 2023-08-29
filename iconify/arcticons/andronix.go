package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Andronix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.669 18.816a10.669 10.669 0 1 0-21.338 0m0 5.627v14.093c0 3.634 2.57 4.964 4.432 4.964s3.28-.532 3.28-3.723s-1.241-7.534-7.712-15.334Zm10.747 3.014c1.364-.532 2.438-3.235 2.438-3.235a18.043 18.043 0 0 0-2.438-.235a18.043 18.043 0 0 0-2.437.235s1.074 2.703 2.437 3.235Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.808 18.834a1.09 1.09 0 0 1-1.09-1.089a1.073 1.073 0 0 1 .074-.366a1.855 1.855 0 1 0 1.382 1.382a1.074 1.074 0 0 1-.366.073Zm11.232 0a1.09 1.09 0 0 1-1.089-1.089a1.07 1.07 0 0 1 .074-.366a1.855 1.855 0 1 0 1.382 1.382a1.073 1.073 0 0 1-.366.073Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.669 18.913a5.334 5.334 0 1 0-10.669 0a5.334 5.334 0 1 0-10.669 0m21.338 5.53v14.093c0 3.634-2.57 4.964-4.432 4.964s-3.28-.532-3.28-3.723s1.241-7.534 7.712-15.334ZM23.052 8.189A6.001 6.001 0 0 0 22.86 4.5m2.991 1.363a10.371 10.371 0 0 0-1.262 2.3"/>`),
		g.Group(children),
	)
}