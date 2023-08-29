package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f2302id3)"><path fill="url(#f2302id0)" d="M11.74 2h8.52c.99 0 1.93.39 2.63 1.09l6.02 6.02c.7.7 1.09 1.64 1.09 2.63v8.52c0 .99-.39 1.93-1.09 2.63l-6.02 6.02c-.7.7-1.64 1.09-2.63 1.09h-8.52c-.99 0-1.93-.39-2.63-1.09l-6.02-6.02c-.7-.7-1.09-1.64-1.09-2.63v-8.52c0-.99.39-1.93 1.09-2.63l6.02-6.02c.7-.7 1.64-1.09 2.63-1.09Z"/><path fill="url(#f2302id1)" d="M11.74 2h8.52c.99 0 1.93.39 2.63 1.09l6.02 6.02c.7.7 1.09 1.64 1.09 2.63v8.52c0 .99-.39 1.93-1.09 2.63l-6.02 6.02c-.7.7-1.64 1.09-2.63 1.09h-8.52c-.99 0-1.93-.39-2.63-1.09l-6.02-6.02c-.7-.7-1.09-1.64-1.09-2.63v-8.52c0-.99.39-1.93 1.09-2.63l6.02-6.02c.7-.7 1.64-1.09 2.63-1.09Z"/></g><path fill="url(#f2302id2)" d="M12.57 4h6.86c.99 0 1.93.39 2.63 1.09l4.85 4.85c.7.7 1.09 1.64 1.09 2.63v6.86c0 .99-.39 1.93-1.09 2.63l-4.85 4.85c-.7.7-1.64 1.09-2.63 1.09h-6.86c-.99 0-1.93-.39-2.63-1.09l-4.85-4.85c-.7-.7-1.09-1.64-1.09-2.63v-6.86c0-.99.39-1.93 1.09-2.63l4.85-4.85c.7-.7 1.64-1.09 2.63-1.09Z"/><defs><linearGradient id="f2302id0" x1="27.86" x2="8.851" y1="6.975" y2="27.592" gradientUnits="userSpaceOnUse"><stop stop-color="#FBF9FC"/><stop offset=".501" stop-color="#D5C9DD"/><stop offset="1" stop-color="#B9B3BC"/></linearGradient><linearGradient id="f2302id1" x1="16" x2="16" y1="31.578" y2="27.653" gradientUnits="userSpaceOnUse"><stop stop-color="#8C749B"/><stop offset="1" stop-color="#8C749B" stop-opacity="0"/></linearGradient><linearGradient id="f2302id2" x1="29.118" x2="3.998" y1="16" y2="16" gradientUnits="userSpaceOnUse"><stop stop-color="#E53146"/><stop offset="1" stop-color="#CC307A"/></linearGradient><filter id="f2302id3" width="28.75" height="28" x="2" y="2" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".75"/><feGaussianBlur stdDeviation=".75"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.52549 0 0 0 0 0.513726 0 0 0 0 0.537255 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18_12046"/></filter></defs></g>`),
		g.Group(children),
	)
}