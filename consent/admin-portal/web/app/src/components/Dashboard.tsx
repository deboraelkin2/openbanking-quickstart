import React, { useEffect, useState } from "react";
import { Route } from "react-router-dom";
import { Container, makeStyles, Theme, Typography } from "@material-ui/core";
import { useHistory } from "react-router";

import PageToolbar from "./PageToolbar";
import Progress from "./Progress";
import { api } from "../api/api";
import noAccountEmptyState from "./no-accounts-empty-state.svg";
import ClientsList from "./ClientsList";
import Subheader from "./Subheader";
import CustomTabs from "./CustomTabs";
import SearchInput from "./SearchInput";

const useStyles = makeStyles((theme: Theme) => ({
  subtitle: {
    ...theme.custom.body1,
  },
}));

export const searchTabs = (history) => [
  {
    key: "account",
    label: "Account",
    content: (
      <div>
        <SearchInput
          placeholder="Search by account number"
          onSearch={(v) => {
            // history.push(`/accounts/${v}`)
          }}
        />
      </div>
    ),
  },
];

interface PropTypes {
  authorizationServerURL?: string;
  authorizationServerId?: string;
  tenantId?: string;
}

export default function Dashboard({
  authorizationServerURL,
  authorizationServerId,
  tenantId,
}: PropTypes) {
  const [isProgress, setProgress] = useState(true);
  const [clients, setClients] = useState<any>([]);
  const classes = useStyles();
  const history = useHistory();

  useEffect(() => {
    setProgress(true);
    api
      .getClients()
      .then((res) => setClients(res.clients || []))
      .catch((err) => console.log(err))
      .finally(() => setProgress(false));
  }, []);

  // const handleRevokeConsent = (id: string) => {
  //   setProgress(true);
  //   api
  //     .deleteConsent({ id })
  //     .then(api.getClients)
  //     .then((res) => setClients(res.clients || []))
  //     .catch((err) => console.log(err))
  //     .finally(() => setProgress(false));
  // };

  const handleRevokeClient = (id: string) => {
    setProgress(true);
    api
      .deleteClient({ id })
      .then(api.getClients)
      .then((res) => setClients(res.clients || []))
      .catch((err) => console.log(err))
      .finally(() => setProgress(false));
  };

  return (
    <div>
      <PageToolbar
        authorizationServerURL={authorizationServerURL}
        authorizationServerId={authorizationServerId}
        tenantId={tenantId}
      />
      <div style={{ position: "relative" }}>
        {isProgress && <Progress />}
        {!isProgress && (
          <>
            {clients.length === 0 && (
              <div style={{ textAlign: "center", marginTop: 128 }}>
                <Typography variant="h3" style={{ color: "#626576" }}>
                  No authorized 3rd party Applications
                </Typography>
                <img
                  src={noAccountEmptyState}
                  style={{ marginTop: 64 }}
                  alt="empty state"
                />
              </div>
            )}
            {clients.length > 0 && (
              <>
                <Subheader title="Customer Consent Portal">
                  <div className={classes.subtitle}>
                    Search for consents and manage data access on behalf of bank
                    members
                  </div>
                  <div style={{ marginTop: 32 }}>
                    <CustomTabs tabs={searchTabs(history)} />
                  </div>
                </Subheader>
                <Container>
                  <Route
                    exact
                    path="/"
                    render={() => (
                      <ClientsList
                        clients={clients}
                        onRevokeClient={handleRevokeClient}
                      />
                    )}
                  />
                </Container>
              </>
            )}
          </>
        )}
      </div>
    </div>
  );
}
