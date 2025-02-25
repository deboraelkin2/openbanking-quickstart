import { TppAuthenticatedPage } from "../../pages/tpp/TppAuthenticatedPage";
import { TppIntentPage } from "../../pages/tpp/TppIntentPage";
import { TppLoginPage } from "../../pages/tpp/TppLoginPage";
import { AcpLoginPage } from "../../pages/acp/AcpLoginPage";
import { AccountConsentPage } from "../../pages/consent/AccountConsentPage";
import { TppErrorPage } from "../../pages/TppErrorPage";
import { ErrorPage } from "../../pages/ErrorPage";

describe(`Tpp technical app`, () => {
  const tppAuthenticatedPage: TppAuthenticatedPage = new TppAuthenticatedPage();
  const tppIntentPage: TppIntentPage = new TppIntentPage();
  const tppLoginPage: TppLoginPage = new TppLoginPage();
  const acpLoginPage: AcpLoginPage = new AcpLoginPage();
  const accountConsentPage: AccountConsentPage = new AccountConsentPage();
  const tppErrorPage: TppErrorPage = new TppErrorPage();
  const errorPage: ErrorPage = new ErrorPage();

  const accountsReadPermission: string = `ACCOUNTS_READ`;
  const accountsOverdraftLimitsReadPermission: string = `ACCOUNTS_OVERDRAFT_LIMITS_READ`;
  const resourcesReadPermission: string = `RESOURCES_READ`;

  beforeEach(() => {
    tppLoginPage.visit();
  });

  [
    [
      accountsReadPermission,
      accountsOverdraftLimitsReadPermission,
      resourcesReadPermission,
    ],
    // [accountsReadPermission, resourcesReadPermission], (example) 1 or 2 permissions selected - UI error page improvements AUT-5845
    [], // none permissions selected - UI error page improvements AUT-5845
  ].forEach((permissions) => {
    it(`Happy path with permissions: ${permissions}`, () => {
      tppLoginPage.checkAccountsReadPermission(
        permissions.includes(accountsReadPermission)
      );
      tppLoginPage.checkAccountsOverdraftLimitsReadPermission(
        permissions.includes(accountsOverdraftLimitsReadPermission)
      );
      tppLoginPage.checkResourcesReadPermission(
        permissions.includes(resourcesReadPermission)
      );
      tppLoginPage.next();
      if (
        !permissions.includes(accountsReadPermission) ||
        !permissions.includes(accountsOverdraftLimitsReadPermission) ||
        !permissions.includes(resourcesReadPermission)
      ) {
        errorPage.assertError(`failed to register consent`);
        return;
      }
      tppIntentPage.login();

      acpLoginPage.assertThatModalIsDisplayed();
      acpLoginPage.loginWithMfaOption();

      accountConsentPage.expandPermissions();
      accountConsentPage.assertPermissionsDetails(
        "Purpose for sharing data",
        "To uncover insights that can improve your financial well being."
      );
      accountConsentPage.clickAgree();
      tppAuthenticatedPage.assertSuccess();
    });
  });

  it(`Cancel on ACP login`, () => {
    tppLoginPage.next();
    tppIntentPage.login();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.cancelLogin();
    
    tppErrorPage.assertThatCancelLoginErrorPageIsDisplayed(
      `access denied`,
      `The user rejected the authentication`
    );
  });

  it(`Cancel on consent`, () => {
    tppLoginPage.next();
    tppIntentPage.login();

    acpLoginPage.assertThatModalIsDisplayed();
    acpLoginPage.loginWithMfaOption();

    accountConsentPage.clickCancel();

    tppErrorPage.assertThatRejectConsentErrorPageIsDisplayed(
      `rejected`,
      `The user rejected the authentication.`,
      `consent_rejected`
    );
  });
});
