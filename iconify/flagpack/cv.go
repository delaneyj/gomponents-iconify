package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCv0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCv2)"><use href="#flagpackCv0"/><path fill="#4141DB" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackCv1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackCv1)"><path fill="#F90000" stroke="#F7FCFF" stroke-width="2" d="M0 13h-1v4h34v-4H0Z"/><g filter="url(#flagpackCv3)"><path fill="#FFDE00" fill-rule="evenodd" d="m9.796 10.26l.726-.51l.702.51l-.218-.906l.564-.574h-.718l-.331-.74l-.283.74h-.842l.645.574l-.245.905Zm3 1.6l.726-.51l.702.51l-.218-.906l.564-.574h-.718l-.331-.74l-.283.74h-.842l.645.574l-.245.905Zm3.126 1.89l-.726.51l.245-.906l-.645-.574h.842l.283-.74l.33.74h.719l-.564.574l.218.905l-.702-.508Zm-.726 4.51l.726-.51l.702.51l-.218-.906l.564-.574h-.718l-.331-.74l-.283.74h-.842l.645.574l-.245.905Zm-1.674 1.89l-.726.51l.245-.906l-.645-.574h.842l.283-.74l.33.74h.719l-.564.574l.218.905l-.702-.508Zm-3.726 2.11l.726-.51l.702.51l-.218-.906l.564-.574h-.718l-.331-.74l-.283.74h-.842l.645.574l-.245.905Zm-2.274-2.11l-.726.51l.245-.906l-.645-.574h.842l.283-.74l.33.74h.719l-.564.574l.218.905l-.702-.508Zm-3.126-1.89l.726-.51l.702.51l-.218-.906l.564-.574h-.718l-.331-.74l-.283.74h-.842l.645.574l-.245.905Zm.726-4.51l-.726.51l.245-.906l-.645-.574h.842l.283-.74l.33.74h.719l-.564.574l.218.905l-.702-.508Zm1.674-1.89l.726-.51l.702.51l-.218-.906l.564-.574h-.718l-.331-.74l-.283.74h-.842l.645.574l-.245.905Z" clip-rule="evenodd"/></g></g></g><defs><clipPath id="flagpackCv2"><use href="#flagpackCv0"/></clipPath><filter id="flagpackCv3" width="20.974" height="22.219" x="-.004" y="4.041" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset/><feGaussianBlur stdDeviation="2"/><feColorMatrix values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.24 0"/><feBlend in2="BackgroundImageFix" result="effect1_dropShadow_270_67487"/><feBlend in="SourceGraphic" in2="effect1_dropShadow_270_67487" result="shape"/></filter></defs></g>`),
		g.Group(children),
	)
}