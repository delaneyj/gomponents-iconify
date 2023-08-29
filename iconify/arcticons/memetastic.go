package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Memetastic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.07 42.5s-1-2.39-.85-3.15a9.86 9.86 0 0 0-.51-3.35c-.17-.38-1-3.91-1.09-4.42s-.34-5.09 0-5.51a5.86 5.86 0 0 0 1.07-2a5.25 5.25 0 0 1 .4-2.65c.46-.67 6-6.85 6.86-7.19s1.93-.5 2.23-.71A24.63 24.63 0 0 1 26 12.27c1.72 0 6.52 2.1 7.7 6.35s3.42 5.38 3.8 6.14s.51 7.49.51 9.38s.71 2.86.5 3.53s-1.62 4.83-1.62 4.83"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.36 34.77a3.74 3.74 0 0 0 4.25 2.61c2.48-.47 2.73-1.73 2.65-3.07a16.31 16.31 0 0 0-.17-1.81m-1.56-14.45a3.29 3.29 0 0 1 3.33.93c1.13 1.26 3.74 4.45 2.69 6s-4.25 2.48-5.72 1.51s-3.83-2.9-3.49-4.16s2.44-4 3.19-4.28Zm-8.99 3.81c1.21.7 3.11 1.6 2.6 3s-2.39 4.2-3.53 4.58s-3.28.38-4-.76s-.88-4.37-.37-5s3.25-3.01 5.3-1.82Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.57 23.41c.68.16 1.71.52 1.61 1.77c-.07.88-1.83 1.37-2.17 1.33M30.86 19a1.49 1.49 0 0 0-1.76 1.78a1.73 1.73 0 0 0 2 .8a1.47 1.47 0 0 0 .9-1.26"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}