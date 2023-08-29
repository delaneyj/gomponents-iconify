package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musicbrainz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#8967AA" d="M34 9.146L12 21.75v27.5l22 12.604V9.146Z"/><path fill="#F4AA41" d="M38 9.146L60 21.75v27.5L38 61.854V9.146Z"/><path fill="#B399C8" d="M18 26.76L23 24l5 2.76V47l-5-2.76L18 47V26.76Z"/><path d="m12 21.75l-.497-.868a1 1 0 0 0-.503.868h1ZM34 9.146h1a1 1 0 0 0-1.497-.868l.497.868ZM12 49.25h-1a1 1 0 0 0 .503.868L12 49.25Zm22 12.604l-.497.868A1 1 0 0 0 35 61.854h-1ZM60 49.25l.497.868A1 1 0 0 0 61 49.25h-1ZM38 61.854h-1a1 1 0 0 0 1.497.868L38 61.854ZM60 21.75h1a1 1 0 0 0-.503-.868L60 21.75ZM38 9.146l.497-.868A1 1 0 0 0 37 9.146h1ZM12.497 22.618l22-12.604l-.994-1.736l-22 12.604l.994 1.736ZM13 49.25v-27.5h-2v27.5h2Zm21.497 11.736l-22-12.604l-.994 1.736l22 12.604l.994-1.736Zm.503.868V9.146h-2v52.708h2Zm24.503-13.472l-22 12.605l.994 1.735l22-12.604l-.994-1.736ZM59 21.75v27.5h2v-27.5h-2ZM37.503 10.014l22 12.604l.994-1.736l-22-12.604l-.994 1.736ZM37 9.146v52.708h2V9.146h-2Z"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M18 26.76L23 24l5 2.76V47l-5-2.76L18 47V26.76Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-width="2" d="M31 16.055A9.004 9.004 0 0 0 23.055 24"/><circle cx="47" cy="28" r="2" fill="none" stroke="#000" stroke-width="2"/><circle cx="47" cy="42" r="2" fill="none" stroke="#000" stroke-width="2"/><circle cx="47" cy="20" r="2" fill="none" stroke="#000" stroke-width="2"/><circle cx="55" cy="41" r="2" fill="none" stroke="#000" stroke-width="2"/><circle cx="55" cy="33" r="2" fill="none" stroke="#000" stroke-width="2"/><circle cx="55" cy="25" r="2" fill="none" stroke="#000" stroke-width="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M38 30h2l4 4h4l5 5"/><path fill="none" stroke="#000" stroke-width="2" d="M50.5 37.5L53 35m2-4v-4"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M38 16h4l3 3"/><circle r="2" fill="none" stroke="#000" stroke-width="2" transform="matrix(1 0 0 -1 47 50)"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M38 54h4l3-3"/><path fill="none" stroke="#000" stroke-width="2" d="M47 22v4m0 18v4"/>`),
		g.Group(children),
	)
}