```bash
$ go test -bench . -cpu 1,2,4
```

```text
PASS
BenchmarkNfntResizeNearestNeighbor	      10	 165852654 ns/op
BenchmarkNfntResizeNearestNeighbor-2	      20	  85096788 ns/op
BenchmarkNfntResizeNearestNeighbor-4	      50	  51456224 ns/op
BenchmarkNfntResizeBilinear	      10	 215553709 ns/op
BenchmarkNfntResizeBilinear-2	      20	 109390764 ns/op
BenchmarkNfntResizeBilinear-4	      20	  73458031 ns/op
BenchmarkNfntResizeBicubic	       5	 340822474 ns/op
BenchmarkNfntResizeBicubic-2	      10	 172641220 ns/op
BenchmarkNfntResizeBicubic-4	      20	 113354571 ns/op
BenchmarkNfntResizeMitchellNetravali	       5	 331845788 ns/op
BenchmarkNfntResizeMitchellNetravali-2	      10	 170286081 ns/op
BenchmarkNfntResizeMitchellNetravali-4	      20	 111036403 ns/op
BenchmarkNfntResizeLanczos2	       5	 334637470 ns/op
BenchmarkNfntResizeLanczos2-2	      10	 170975450 ns/op
BenchmarkNfntResizeLanczos2-4	      20	 113139111 ns/op
BenchmarkNfntResizeLanczos3	       5	 453353474 ns/op
BenchmarkNfntResizeLanczos3-2	      10	 231613014 ns/op
BenchmarkNfntResizeLanczos3-4	      10	 154705916 ns/op
BenchmarkGiftResizeNearestNeighbor	      50	  32781836 ns/op
BenchmarkGiftResizeNearestNeighbor-2	     100	  16717827 ns/op
BenchmarkGiftResizeNearestNeighbor-4	     100	  11116779 ns/op
BenchmarkGiftResizeLinear	       5	 400236762 ns/op
BenchmarkGiftResizeLinear-2	      10	 205392892 ns/op
BenchmarkGiftResizeLinear-4	      10	 139582966 ns/op
BenchmarkGiftResizeBox	       5	 331759408 ns/op
BenchmarkGiftResizeBox-2	      10	 171398995 ns/op
BenchmarkGiftResizeBox-4	      20	 120437284 ns/op
BenchmarkGiftResizeCubic	       2	 504086344 ns/op
BenchmarkGiftResizeCubic-2	       5	 260958743 ns/op
BenchmarkGiftResizeCubic-4	      10	 173766370 ns/op
BenchmarkGiftResizeLanczos	       2	 615086038 ns/op
BenchmarkGiftResizeLanczos-2	       5	 320190437 ns/op
BenchmarkGiftResizeLanczos-4	      10	 208883762 ns/op
BenchmarkMoustachioResample	      50	  46488351 ns/op
BenchmarkMoustachioResample-2	      50	  46520872 ns/op
BenchmarkMoustachioResample-4	      50	  46539056 ns/op
BenchmarkMoustachioResize	       2	 590352132 ns/op
BenchmarkMoustachioResize-2	       2	 589678547 ns/op
BenchmarkMoustachioResize-4	       2	 589989167 ns/op
BenchmarkRezConvertBilinear	     100	  25696497 ns/op
BenchmarkRezConvertBilinear-2	     100	  15516646 ns/op
BenchmarkRezConvertBilinear-4	     100	  13002075 ns/op
BenchmarkRezConvertBicubic	      50	  44685808 ns/op
BenchmarkRezConvertBicubic-2	     100	  24578122 ns/op
BenchmarkRezConvertBicubic-4	     100	  21533706 ns/op
BenchmarkRezConvertLanczos2	      50	  47539609 ns/op
BenchmarkRezConvertLanczos2-2	     100	  28356712 ns/op
BenchmarkRezConvertLanczos2-4	     100	  22716406 ns/op
BenchmarkRezConvertLanczos3	      20	  71581173 ns/op
BenchmarkRezConvertLanczos3-2	      50	  39953376 ns/op
BenchmarkRezConvertLanczos3-4	      50	  33615933 ns/op
BenchmarkGraphicsScale	      20	  97321436 ns/op
BenchmarkGraphicsScale-2	      20	  97361626 ns/op
BenchmarkGraphicsScale-4	      20	  97312046 ns/op
BenchmarkImagingResizeNearestNeighbor	      50	  61220687 ns/op
BenchmarkImagingResizeNearestNeighbor-2	      50	  50122711 ns/op
BenchmarkImagingResizeNearestNeighbor-4	      50	  45306938 ns/op
BenchmarkImagingResizeBox	       5	 311383842 ns/op
BenchmarkImagingResizeBox-2	      10	 183574203 ns/op
BenchmarkImagingResizeBox-4	      20	 125230628 ns/op
BenchmarkImagingResizeLinear	       5	 456764844 ns/op
BenchmarkImagingResizeLinear-2	       5	 254993264 ns/op
BenchmarkImagingResizeLinear-4	      10	 164408588 ns/op
BenchmarkImagingResizeHermite	       5	 456273468 ns/op
BenchmarkImagingResizeHermite-2	       5	 255466033 ns/op
BenchmarkImagingResizeHermite-4	      10	 164340500 ns/op
BenchmarkImagingResizeMitchellNetravali	       2	 670244030 ns/op
BenchmarkImagingResizeMitchellNetravali-2	       5	 365345733 ns/op
BenchmarkImagingResizeMitchellNetravali-4	      10	 226548494 ns/op
BenchmarkImagingResizeCatmullRom	       2	 673925993 ns/op
BenchmarkImagingResizeCatmullRom-2	       5	 365456501 ns/op
BenchmarkImagingResizeCatmullRom-4	      10	 226707237 ns/op
BenchmarkImagingResizeBSpline	       2	 670560237 ns/op
BenchmarkImagingResizeBSpline-2	       5	 371825914 ns/op
BenchmarkImagingResizeBSpline-4	      10	 226981175 ns/op
BenchmarkImagingResizeGaussian	       2	 672459884 ns/op
BenchmarkImagingResizeGaussian-2	       5	 364735264 ns/op
BenchmarkImagingResizeGaussian-4	      10	 226606344 ns/op
BenchmarkImagingResizeBartlett	       2	 894683465 ns/op
BenchmarkImagingResizeBartlett-2	       5	 476563647 ns/op
BenchmarkImagingResizeBartlett-4	       5	 288349796 ns/op
BenchmarkImagingResizeLanczos	       2	 892661874 ns/op
BenchmarkImagingResizeLanczos-2	       5	 476484558 ns/op
BenchmarkImagingResizeLanczos-4	       5	 292293105 ns/op
BenchmarkImagingResizeHann	       2	 894416916 ns/op
BenchmarkImagingResizeHann-2	       5	 477643560 ns/op
BenchmarkImagingResizeHann-4	       5	 289540324 ns/op
BenchmarkImagingResizeHamming	       2	 896447935 ns/op
BenchmarkImagingResizeHamming-2	       5	 480581783 ns/op
BenchmarkImagingResizeHamming-4	       5	 288107405 ns/op
BenchmarkImagingResizeWelch	       2	 892771245 ns/op
BenchmarkImagingResizeWelch-2	       5	 476364780 ns/op
BenchmarkImagingResizeWelch-4	       5	 290368796 ns/op
BenchmarkImagingResizeCosine	       2	 896756145 ns/op
BenchmarkImagingResizeCosine-2	       5	 477968252 ns/op
BenchmarkImagingResizeCosine-4	       5	 291363587 ns/op
ok  	github.com/oov/resize-test	207.854s
```