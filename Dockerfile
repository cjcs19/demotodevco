# iron/go is the alpine image with only ca-certificates added
FROM iron/go

WORKDIR /

# Now just add the binary
ADD demotodevco /

EXPOSE 3001
ENTRYPOINT ["./demotodevco"]


