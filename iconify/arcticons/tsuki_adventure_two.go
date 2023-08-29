package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TsukiAdventureTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.096 30.99c-.16-1.486.764-1.628 1.394-1.743c-.4-1.07.27-1.485 1.008-1.65m-28.636 3.668c4.196-.804 17.073-.124 26.035-.145l2.603-.015m-34.682-.019L5.5 31.55m4.976-11.869c-1.614-.717-3.433 2.163-4.976 4.914m28.663 1.123c.722.6.86.983 1.069 1.582"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.612 27.44c1.103 1.47 2.406 2.556 4.613 2.058c.97-.22 2.03.888 2.815 1.53l1.913-2.26c-.585-.167-1.286-.238-2.513-1.066c-.383-.258-.806-.425-1.639-.278c-.938.166-1.734.046-2.29-.938"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.556 27.727c3.537.31 7.247.267 11.243-1.969c.84-.586.14-1.868-.323-2.363m-3.155-9.038c.772-1.965 5.15-4.614 5.799-4.17m-18.713 10.03l.87-.074m3.698-1.46v.646"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.805 15.823c2.348-1.584 6.877-4.763 6.257-7.267c-.944-2.021-9.764.674-10.8 5.537c-2.01-1.713-10.703.837-10.897 5.954c-1.09 4.549 6.456 8.053 14.743 4.491c1.438-.871 3.69-5.099.698-8.715h-.001ZM15.207 25.89c-.926-.314-.963 1.561-1.321 2.598m2.697.288l.057 2.078"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.728 29.916c-5.045-1.126-4.619-3.857-4.131-4.82c1.786-2.225 6.753 2.687 7.232 2.987m12.61 6.84c-.432 1.726-3.528 1.275-4.058 1.021c-.207-1.047 1.06-1.325 2.2-1.432c1.435-.275 1.9.246 1.858.41v.001Zm2.08-2.022c.512 2.86.152 6.436.153 9.599m-27.084-9.175c-.568 1.493-.84 6.191-.476 9.175m.559-23.287c-.244 2.757-.681 5.462-.685 8.031m-.506 1.89c-.451.637-1.237 1.707-1.587 2.578c.086 1.533 1.763.71 2.601.23"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.707 24.8c2.308-.37 2.634 8.728.605 7.838c-1.591-.697-2.735-1.523-2.602-3.774c.138-2.052.75-3.7 1.997-4.064Z"/>`),
		g.Group(children),
	)
}