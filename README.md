```bash
$ go test -bench . -cpu 1,2,4 -timeout=30m
```

```text
BenchmarkScalerNfntResizeNearestNeighbor       	      10	 133409240 ns/op
BenchmarkScalerNfntResizeNearestNeighbor-2     	      10	 128697340 ns/op
BenchmarkScalerNfntResizeNearestNeighbor-4     	      20	  94040120 ns/op
BenchmarkScalerNfntResizeBilinear              	      10	 199757090 ns/op
BenchmarkScalerNfntResizeBilinear-2            	      10	 174207250 ns/op
BenchmarkScalerNfntResizeBilinear-4            	      10	 128857460 ns/op
BenchmarkScalerNfntResizeBicubic               	       5	 290346060 ns/op
BenchmarkScalerNfntResizeBicubic-2             	       5	 251394520 ns/op
BenchmarkScalerNfntResizeBicubic-4             	      10	 158803820 ns/op
BenchmarkScalerNfntResizeMitchellNetravali     	       5	 310188120 ns/op
BenchmarkScalerNfntResizeMitchellNetravali-2   	       5	 231997940 ns/op
BenchmarkScalerNfntResizeMitchellNetravali-4   	      10	 161540130 ns/op
BenchmarkScalerNfntResizeLanczos2              	       5	 292149080 ns/op
BenchmarkScalerNfntResizeLanczos2-2            	       5	 271326380 ns/op
BenchmarkScalerNfntResizeLanczos2-4            	      10	 160008210 ns/op
BenchmarkScalerNfntResizeLanczos3              	       3	 406938800 ns/op
BenchmarkScalerNfntResizeLanczos3-2            	       3	 341369966 ns/op
BenchmarkScalerNfntResizeLanczos3-4            	       5	 206273940 ns/op
BenchmarkScalerGiftResizeNearestNeighbor       	     500	   3443350 ns/op
BenchmarkScalerGiftResizeNearestNeighbor-2     	    1000	   1720945 ns/op
BenchmarkScalerGiftResizeNearestNeighbor-4     	    2000	    936905 ns/op
BenchmarkScalerGiftResizeLinear                	       3	 427990700 ns/op
BenchmarkScalerGiftResizeLinear-2              	       5	 227551780 ns/op
BenchmarkScalerGiftResizeLinear-4              	      10	 119463020 ns/op
BenchmarkScalerGiftResizeBox                   	       3	 381772000 ns/op
BenchmarkScalerGiftResizeBox-2                 	      10	 202467090 ns/op
BenchmarkScalerGiftResizeBox-4                 	      10	 102325720 ns/op
BenchmarkScalerGiftResizeCubic                 	       2	 530349200 ns/op
BenchmarkScalerGiftResizeCubic-2               	       5	 270615820 ns/op
BenchmarkScalerGiftResizeCubic-4               	      10	 143393100 ns/op
BenchmarkScalerGiftResizeLanczos               	       2	 631721450 ns/op
BenchmarkScalerGiftResizeLanczos-2             	       5	 322195000 ns/op
BenchmarkScalerGiftResizeLanczos-4             	      10	 170011410 ns/op
BenchmarkScalerMoustachioResample              	     200	   6116850 ns/op
BenchmarkScalerMoustachioResample-2            	     300	   5890260 ns/op
BenchmarkScalerMoustachioResample-4            	     300	   5979730 ns/op
BenchmarkScalerMoustachioResize                	       2	 634725800 ns/op
BenchmarkScalerMoustachioResize-2              	       2	 633222300 ns/op
BenchmarkScalerMoustachioResize-4              	       2	 633472850 ns/op
BenchmarkScalerRezConvertBilinear              	     100	  10831650 ns/op
BenchmarkScalerRezConvertBilinear-2            	     200	   6266599 ns/op
BenchmarkScalerRezConvertBilinear-4            	     300	   4258860 ns/op
BenchmarkScalerRezConvertBicubic               	     100	  20972323 ns/op
BenchmarkScalerRezConvertBicubic-2             	     100	  11113095 ns/op
BenchmarkScalerRezConvertBicubic-4             	     200	   6660372 ns/op
BenchmarkScalerRezConvertLanczos2              	     100	  21715732 ns/op
BenchmarkScalerRezConvertLanczos2-2            	     100	  12305165 ns/op
BenchmarkScalerRezConvertLanczos2-4            	     200	   7392249 ns/op
BenchmarkScalerRezConvertLanczos3              	      30	  43750133 ns/op
BenchmarkScalerRezConvertLanczos3-2            	      50	  25150666 ns/op
BenchmarkScalerRezConvertLanczos3-4            	     100	  13752165 ns/op
BenchmarkScalerGraphicsScale                   	      50	  29664494 ns/op
BenchmarkScalerGraphicsScale-2                 	      50	  27139324 ns/op
BenchmarkScalerGraphicsScale-4                 	      50	  27593602 ns/op
BenchmarkScalerImagingResizeNearestNeighbor    	       5	 218435820 ns/op
BenchmarkScalerImagingResizeNearestNeighbor-2  	      10	 123696660 ns/op
BenchmarkScalerImagingResizeNearestNeighbor-4  	      20	  64676835 ns/op
BenchmarkScalerImagingResizeBox                	       3	 337689366 ns/op
BenchmarkScalerImagingResizeBox-2              	      10	 176534920 ns/op
BenchmarkScalerImagingResizeBox-4              	      20	  92589170 ns/op
BenchmarkScalerImagingResizeLinear             	       3	 400786400 ns/op
BenchmarkScalerImagingResizeLinear-2           	       5	 219757960 ns/op
BenchmarkScalerImagingResizeLinear-4           	      10	 118457560 ns/op
BenchmarkScalerImagingResizeHermite            	       3	 399280466 ns/op
BenchmarkScalerImagingResizeHermite-2          	       5	 218574380 ns/op
BenchmarkScalerImagingResizeHermite-4          	      10	 119299340 ns/op
BenchmarkScalerImagingResizeMitchellNetravali  	       2	 581913300 ns/op
BenchmarkScalerImagingResizeMitchellNetravali-2	       5	 295088020 ns/op
BenchmarkScalerImagingResizeMitchellNetravali-4	      10	 154695930 ns/op
BenchmarkScalerImagingResizeCatmullRom         	       2	 555636650 ns/op
BenchmarkScalerImagingResizeCatmullRom-2       	       5	 297923980 ns/op
BenchmarkScalerImagingResizeCatmullRom-4       	      10	 157126390 ns/op
BenchmarkScalerImagingResizeBSpline            	       2	 521551550 ns/op
BenchmarkScalerImagingResizeBSpline-2          	       5	 276258460 ns/op
BenchmarkScalerImagingResizeBSpline-4          	      10	 145574330 ns/op
BenchmarkScalerImagingResizeGaussian           	       2	 514803950 ns/op
BenchmarkScalerImagingResizeGaussian-2         	       5	 279173680 ns/op
BenchmarkScalerImagingResizeGaussian-4         	      10	 144925920 ns/op
BenchmarkScalerImagingResizeBartlett           	       2	 699074550 ns/op
BenchmarkScalerImagingResizeBartlett-2         	       3	 379548200 ns/op
BenchmarkScalerImagingResizeBartlett-4         	       5	 200811680 ns/op
BenchmarkScalerImagingResizeLanczos            	       2	 697354550 ns/op
BenchmarkScalerImagingResizeLanczos-2          	       3	 370240133 ns/op
BenchmarkScalerImagingResizeLanczos-4          	      10	 198986860 ns/op
BenchmarkScalerImagingResizeHann               	       2	 665523150 ns/op
BenchmarkScalerImagingResizeHann-2             	       3	 358769966 ns/op
BenchmarkScalerImagingResizeHann-4             	      10	 185822740 ns/op
BenchmarkScalerImagingResizeHamming            	       2	 699074600 ns/op
BenchmarkScalerImagingResizeHamming-2          	       3	 376213500 ns/op
BenchmarkScalerImagingResizeHamming-4          	      10	 200558660 ns/op
BenchmarkScalerImagingResizeWelch              	       2	 701828850 ns/op
BenchmarkScalerImagingResizeWelch-2            	       3	 377313633 ns/op
BenchmarkScalerImagingResizeWelch-4            	       5	 204428540 ns/op
BenchmarkScalerImagingResizeCosine             	       2	 698836600 ns/op
BenchmarkScalerImagingResizeCosine-2           	       3	 382106600 ns/op
BenchmarkScalerImagingResizeCosine-4           	      10	 197110290 ns/op
BenchmarkScalerDrawResizeNearestNeighbor       	    1000	   1508819 ns/op
BenchmarkScalerDrawResizeNearestNeighbor-2     	    1000	   1522663 ns/op
BenchmarkScalerDrawResizeNearestNeighbor-4     	    1000	   1527720 ns/op
BenchmarkScalerDrawResizeApproxBiLinear        	     500	   3431299 ns/op
BenchmarkScalerDrawResizeApproxBiLinear-2      	     500	   3428701 ns/op
BenchmarkScalerDrawResizeApproxBiLinear-4      	     500	   3474005 ns/op
BenchmarkScalerDrawResizeBiLinear              	       5	 294279920 ns/op
BenchmarkScalerDrawResizeBiLinear-2            	       5	 279034180 ns/op
BenchmarkScalerDrawResizeBiLinear-4            	       5	 284734060 ns/op
BenchmarkScalerDrawResizeCatmullRom            	       2	 530315100 ns/op
BenchmarkScalerDrawResizeCatmullRom-2          	       2	 536575550 ns/op
BenchmarkScalerDrawResizeCatmullRom-4          	       2	 544879750 ns/op
BenchmarkFlowNfntResizeNearestNeighbor         	       2	 556161800 ns/op
BenchmarkFlowNfntResizeNearestNeighbor-2       	       2	 529585400 ns/op
BenchmarkFlowNfntResizeNearestNeighbor-4       	       2	 529580400 ns/op
BenchmarkFlowNfntResizeBilinear                	       2	 622736250 ns/op
BenchmarkFlowNfntResizeBilinear-2              	       2	 571882500 ns/op
BenchmarkFlowNfntResizeBilinear-4              	       2	 537861150 ns/op
BenchmarkFlowNfntResizeBicubic                 	       2	 755160850 ns/op
BenchmarkFlowNfntResizeBicubic-2               	       2	 632236150 ns/op
BenchmarkFlowNfntResizeBicubic-4               	       2	 581530800 ns/op
BenchmarkFlowNfntResizeMitchellNetravali       	       2	 722861150 ns/op
BenchmarkFlowNfntResizeMitchellNetravali-2     	       2	 614397750 ns/op
BenchmarkFlowNfntResizeMitchellNetravali-4     	       2	 556165650 ns/op
BenchmarkFlowNfntResizeLanczos2                	       2	 753408150 ns/op
BenchmarkFlowNfntResizeLanczos2-2              	       2	 638759250 ns/op
BenchmarkFlowNfntResizeLanczos2-4              	       2	 556904000 ns/op
BenchmarkFlowNfntResizeLanczos3                	       2	 813605000 ns/op
BenchmarkFlowNfntResizeLanczos3-2              	       2	 687073550 ns/op
BenchmarkFlowNfntResizeLanczos3-4              	       2	 596423150 ns/op
BenchmarkFlowGiftResizeNearestNeighbor         	       3	 421147366 ns/op
BenchmarkFlowGiftResizeNearestNeighbor-2       	       3	 418111666 ns/op
BenchmarkFlowGiftResizeNearestNeighbor-4       	       3	 406463400 ns/op
BenchmarkFlowGiftResizeLinear                  	       2	 839383500 ns/op
BenchmarkFlowGiftResizeLinear-2                	       2	 652642450 ns/op
BenchmarkFlowGiftResizeLinear-4                	       2	 536574900 ns/op
BenchmarkFlowGiftResizeBox                     	       2	 786977850 ns/op
BenchmarkFlowGiftResizeBox-2                   	       2	 607159100 ns/op
BenchmarkFlowGiftResizeBox-4                   	       2	 517796050 ns/op
BenchmarkFlowGiftResizeCubic                   	       2	 933434900 ns/op
BenchmarkFlowGiftResizeCubic-2                 	       2	 691549450 ns/op
BenchmarkFlowGiftResizeCubic-4                 	       2	 551598650 ns/op
BenchmarkFlowGiftResizeLanczos                 	       1	1031629200 ns/op
BenchmarkFlowGiftResizeLanczos-2               	       2	 733197000 ns/op
BenchmarkFlowGiftResizeLanczos-4               	       2	 581842200 ns/op
BenchmarkFlowMoustachioResample                	       3	 409466566 ns/op
BenchmarkFlowMoustachioResample-2              	       3	 405970600 ns/op
BenchmarkFlowMoustachioResample-4              	       3	 410771733 ns/op
BenchmarkFlowMoustachioResize                  	       1	1086670400 ns/op
BenchmarkFlowMoustachioResize-2                	       1	1087067100 ns/op
BenchmarkFlowMoustachioResize-4                	       1	1027508500 ns/op
BenchmarkFlowRezConvertBilinear                	       3	 440510466 ns/op
BenchmarkFlowRezConvertBilinear-2              	       3	 440949600 ns/op
BenchmarkFlowRezConvertBilinear-4              	       3	 446279266 ns/op
BenchmarkFlowRezConvertBicubic                 	       3	 478749633 ns/op
BenchmarkFlowRezConvertBicubic-2               	       3	 454499400 ns/op
BenchmarkFlowRezConvertBicubic-4               	       3	 454042933 ns/op
BenchmarkFlowRezConvertLanczos2                	       3	 454201233 ns/op
BenchmarkFlowRezConvertLanczos2-2              	       3	 461979633 ns/op
BenchmarkFlowRezConvertLanczos2-4              	       3	 440750500 ns/op
BenchmarkFlowRezConvertLanczos3                	       3	 480568400 ns/op
BenchmarkFlowRezConvertLanczos3-2              	       3	 462544100 ns/op
BenchmarkFlowRezConvertLanczos3-4              	       3	 455314466 ns/op
BenchmarkFlowGraphicsScale                     	       3	 448367300 ns/op
BenchmarkFlowGraphicsScale-2                   	       3	 447521333 ns/op
BenchmarkFlowGraphicsScale-4                   	       3	 439598500 ns/op
BenchmarkFlowImagingResizeNearestNeighbor      	       2	 618701300 ns/op
BenchmarkFlowImagingResizeNearestNeighbor-2    	       2	 518441250 ns/op
BenchmarkFlowImagingResizeNearestNeighbor-4    	       3	 460165666 ns/op
BenchmarkFlowImagingResizeBox                  	       2	 733121300 ns/op
BenchmarkFlowImagingResizeBox-2                	       2	 599671750 ns/op
BenchmarkFlowImagingResizeBox-4                	       2	 513486200 ns/op
BenchmarkFlowImagingResizeLinear               	       2	 812498950 ns/op
BenchmarkFlowImagingResizeLinear-2             	       2	 629682250 ns/op
BenchmarkFlowImagingResizeLinear-4             	       2	 528706800 ns/op
BenchmarkFlowImagingResizeHermite              	       2	 848821050 ns/op
BenchmarkFlowImagingResizeHermite-2            	       2	 653966600 ns/op
BenchmarkFlowImagingResizeHermite-4            	       2	 538374750 ns/op
BenchmarkFlowImagingResizeMitchellNetravali    	       1	1006559500 ns/op
BenchmarkFlowImagingResizeMitchellNetravali-2  	       2	 724572900 ns/op
BenchmarkFlowImagingResizeMitchellNetravali-4  	       2	 592209650 ns/op
BenchmarkFlowImagingResizeCatmullRom           	       2	 956523450 ns/op
BenchmarkFlowImagingResizeCatmullRom-2         	       2	 718855800 ns/op
BenchmarkFlowImagingResizeCatmullRom-4         	       2	 568854950 ns/op
BenchmarkFlowImagingResizeBSpline              	       2	 933186200 ns/op
BenchmarkFlowImagingResizeBSpline-2            	       2	 691146900 ns/op
BenchmarkFlowImagingResizeBSpline-4            	       2	 571569350 ns/op
BenchmarkFlowImagingResizeGaussian             	       2	 933935650 ns/op
BenchmarkFlowImagingResizeGaussian-2           	       2	 691219650 ns/op
BenchmarkFlowImagingResizeGaussian-4           	       2	 558550000 ns/op
BenchmarkFlowImagingResizeBartlett             	       1	1113223500 ns/op
BenchmarkFlowImagingResizeBartlett-2           	       2	 786485150 ns/op
BenchmarkFlowImagingResizeBartlett-4           	       2	 604171650 ns/op
BenchmarkFlowImagingResizeLanczos              	       1	1096218200 ns/op
BenchmarkFlowImagingResizeLanczos-2            	       2	 786442050 ns/op
BenchmarkFlowImagingResizeLanczos-4            	       2	 602395900 ns/op
BenchmarkFlowImagingResizeHann                 	       1	1068638000 ns/op
BenchmarkFlowImagingResizeHann-2               	       2	 762579000 ns/op
BenchmarkFlowImagingResizeHann-4               	       2	 597622050 ns/op
BenchmarkFlowImagingResizeHamming              	       1	1119220400 ns/op
BenchmarkFlowImagingResizeHamming-2            	       2	 811229000 ns/op
BenchmarkFlowImagingResizeHamming-4            	       2	 611193300 ns/op
BenchmarkFlowImagingResizeWelch                	       1	1118231400 ns/op
BenchmarkFlowImagingResizeWelch-2              	       2	1151500400 ns/op
BenchmarkFlowImagingResizeWelch-4              	       2	 797547200 ns/op
BenchmarkFlowImagingResizeCosine               	       1	1177342200 ns/op
BenchmarkFlowImagingResizeCosine-2             	       2	 890057600 ns/op
BenchmarkFlowImagingResizeCosine-4             	       2	 702842650 ns/op
BenchmarkFlowDrawResizeNearestNeighbor         	       3	 412139033 ns/op
BenchmarkFlowDrawResizeNearestNeighbor-2       	       3	 437845933 ns/op
BenchmarkFlowDrawResizeNearestNeighbor-4       	       3	 410486800 ns/op
BenchmarkFlowDrawResizeApproxBiLinear          	       3	 432849433 ns/op
BenchmarkFlowDrawResizeApproxBiLinear-2        	       3	 406971000 ns/op
BenchmarkFlowDrawResizeApproxBiLinear-4        	       3	 427490633 ns/op
BenchmarkFlowDrawResizeBiLinear                	       2	 748662950 ns/op
BenchmarkFlowDrawResizeBiLinear-2              	       2	 706379150 ns/op
BenchmarkFlowDrawResizeBiLinear-4              	       2	 710334000 ns/op
BenchmarkFlowDrawResizeCatmullRom              	       2	 983799800 ns/op
BenchmarkFlowDrawResizeCatmullRom-2            	       1	1010090700 ns/op
BenchmarkFlowDrawResizeCatmullRom-4            	       2	 969999550 ns/op
BenchmarkFlowImageMagickResize                 	       3	 397290900 ns/op
BenchmarkFlowImageMagickResize-2               	       3	 399283400 ns/op
BenchmarkFlowImageMagickResize-4               	       3	 393605000 ns/op
BenchmarkFlowImageMagickResizeWithDefine       	      10	 119081050 ns/op
BenchmarkFlowImageMagickResizeWithDefine-2     	      10	 140666270 ns/op
BenchmarkFlowImageMagickResizeWithDefine-4     	      10	 129523620 ns/op
```
