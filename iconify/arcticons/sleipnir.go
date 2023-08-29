package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleipnir(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.676 20.686a6.794 6.794 0 0 1-4.504 1.758s-2.527 4.98-5.456 4.98a3.938 3.938 0 0 1-3.882-4.175c.147-1.392 5.64-9.045 5.64-9.045s2.233-6.372 8.202-6.372s9.485 4.65 9.485 10.547s-1.941 10.656-1.685 15.307l-1.208 1.867H18.916l-1.245-1.391c0-5.09 3.076-10.73 6.005-13.476Zm6.592 14.868l3.406 4.98l-.513 2.966H16.059l-.695-2.783l3.552-5.163"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.18 7.974c.178-1.094 2.485-3.913 7.172-2.668c-1.72.622-2.05 1.178-2.05 1.178a22.4 22.4 0 0 1 5.456.177a5.22 5.22 0 0 0-1.501 1.647a5.252 5.252 0 0 1 5.126 2.893a3.525 3.525 0 0 0-2.087.293a8.093 8.093 0 0 1 3.479 3.699a21.65 21.65 0 0 0-2.6-.476a9.79 9.79 0 0 1 1.794 2.05c.183.586 1.245 1.685 1.209 2.857a23.652 23.652 0 0 1-.916 3.992l-1.794-1.063s1.575 2.857 1.099 3.955s-2.27.293-2.27.293s1.867 2.417 1.245 3.296s-.476 2.49-1.172 2.747s-2.893.842-2.893.842"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.298 9.193c-1.298-.592-3.1-2.49-2.58-4.693a20.35 20.35 0 0 0 5.463 3.474"/>`),
		g.Group(children),
	)
}