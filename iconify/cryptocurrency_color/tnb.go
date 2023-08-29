package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tnb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#ffc04e"/><path fill="#fff" fill-rule="nonzero" d="M15.598 11.857h.003l-.03.13l-.26 1.227h-.03L13.568 20.5h-2.179l1.747-7.286h-2.932L8.558 20.5H6.162l2.178-9h7.333zm4.141-.357h3.63c3.123.143 2.832 2.214 2.832 2.214H27l-.29.857h-.727c-.217.786-1.67 1.215-1.67 1.215c1.67.214 1.598 1.5 1.598 1.5h.799l-.218.857h-.654c-.167 1.19-1.222 1.768-2.089 2.043a5.35 5.35 0 0 1-1.62.243H17.56zM7.323 13.214H5l.436-1.714h2.323zm9.948 6.286l-1.888-4.429l.799-3.428l1.887 4.214zm5.554-4.429c.26 0 .51-.1.693-.282s.287-.426.287-.682c0-.256-.103-.501-.287-.682s-.433-.282-.693-.282H21.41l-.435 1.928zm-.489 3.715c.651 0 1.179-.496 1.179-1.107c0-.612-.525-1.108-1.179-1.108h-1.702l-.524 2.215z"/></g>`),
		g.Group(children),
	)
}