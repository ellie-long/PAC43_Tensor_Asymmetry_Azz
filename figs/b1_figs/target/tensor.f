      PROGRAM TENSOR


      DO P = 0.,1.,0.01

        A = 2.-SQRT(4.-3.*P**2)

        !WRITE(6,*) P*100.,A*100.

      ENDDO

      DO R=1.0,5.0,0.01

        P = (R**2-1.)/(R**2+R+1.)
        A = (R**2-2*R+1.)/(R**2+R+1.)
        WRITE(6,*) R,P,A
      ENDDO

      RETURN
      END
