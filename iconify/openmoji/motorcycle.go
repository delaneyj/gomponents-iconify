package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Motorcycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<ellipse cx="12.031" cy="48" fill="#9B9B9A" rx="7.031" ry="7"/><ellipse cx="53.031" cy="48" fill="#9B9B9A" rx="8.031" ry="7.995"/><path fill="#EA5A47" d="M12.115 35.12c-1.195-.68-2-1.12-2-1.12s3.784-6.392 9-9c4-2 11-3 11-3v3l-6 4l4 3l8-2l7 6s7-3 10-6v-2l14-3l-1 3s-22 16-28 16l5 6h-20s0-6-1-8c-.536-1.072-3.95-3.169-7-5"/><path fill="#FCEA2B" d="m65.115 30l2 2l-1.698 1.668l-2.542-2.542M15 38l-3-2s4-5 5-5s-2 7-2 7z"/><path fill="#D22F27" d="M28.115 32c0 4.834 4.885 7.667 8.552 7.667c3.064 0 10.114-4.856 12.5-5.334c1.715-.343 1.527-1.893 4.083-2.708c1.585-.506 10.25-4.5 12.75-4.25c1.634.164-2.472 2.251-1.333 1.625c6.667-3.666-25 14.667-25 14.667L37 44l-7.833-3.666l-5.667-5.5v-5.5L28.115 32z"/><path fill="#3F3F3F" d="M39.5 32.5L42 36l8.5-3l2.615-3v-2z"/><path fill="#D0CFCE" d="m40.15 45.388l5.6.581L48.594 45l4.125-2.125l6.906-3.094l.219-1.312L56 34.949l-8.031 5.364l-7.694 3.002L38.87 44l.63 1z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M15.115 37s3-4 2-5"/><ellipse cx="12.031" cy="48" rx="7.031" ry="7"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.115 37c3.05 1.831 6.464 3.928 7 5c1 2 1 9 1 9h20s-.486-6.037-12.281-10c-1.976-.663-11.495-7.224-6.719-12"/><path stroke-linecap="round" stroke-linejoin="round" d="M12.115 35.12c-1.195-.68-2-1.12-2-1.12s3.784-6.392 9-9c4-2 11-3 11-3v3l-6 4l4 3s6.715-3.628 10-1c1.875 1.5 4 5 4 5s8-2 11-5v-3l14-3l-1 3s-22.49 16.167-28.49 16.167"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.115 37s3-4 2-5m1 8l-6.084 8m53.084-18l2 2l-1 1M59 39l-10.885 6H46m7.115-17c-5.28 2.85-10.282 4-14 4"/><path stroke-linecap="round" d="M46.267 49.916A7.032 7.032 0 0 0 53.03 55c3.883 0 7.032-3.134 7.032-7c0-1.66-.58-3.185-1.551-4.385"/></g>`),
		g.Group(children),
	)
}