FROM scratch

ADD openhabskill /openhabskill

ADD configuration /configuration

VOLUME ["/configuration"]

ENTRYPOINT ["/openhabskill"]