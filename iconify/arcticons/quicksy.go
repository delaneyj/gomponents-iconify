package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quicksy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.942 2.504A21.47 21.47 0 0 1 42.932 34.1l.55 2.079l1.929 7.196a1.67 1.67 0 0 1-1.18 2a1.61 1.61 0 0 1-.87 0l-7.196-1.93l-2.109-.46a21.49 21.49 0 1 1-10.115-40.48Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.3 28.8c2.16.104 3.56 4.04 5.9 4c1.558-2.373.941-6.077-1.2-7.3c-2.701-1.544-4.82 1.015-4.7 3.3Zm9.2 7.2s2.4-.045 3.2-.8a3.923 3.923 0 0 0 1-3.2c-.166-1.405-1.251-2.764-.7-4a4.654 4.654 0 0 1 4-2.5a4.654 4.654 0 0 1 4 2.5c.551 1.236-.526 2.67-.7 4a3.88 3.88 0 0 0 1 3.2c.8.755 3.2.8 3.2.8s-4.865 2.5-7.5 2.5s-7.5-2.5-7.5-2.5Zm24.2-7.2c-2.16.104-3.56 4.04-5.9 4c-1.558-2.373-.941-6.077 1.2-7.3c2.701-1.544 4.82 1.015 4.7 3.3Z"/><circle cx="12.75" cy="20.5" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="20.5" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.244" cy="20.502" r="1.999" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}