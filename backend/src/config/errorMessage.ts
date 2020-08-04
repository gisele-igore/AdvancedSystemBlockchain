export namespace ErrorMessage {
  export const NO_COMPAGNIE_ASSURANCE_FOUND = 'noCompagnieAssuranceFound';

  export const COMPAGNIE_ASSURANCE_ALREADY_EXIST = 'compagnieAssuranceAlreadyExist';
 
}

export function getErrorMessage(typeError: string) {
  switch (typeError) {
    case ErrorMessage.COMPAGNIE_ASSURANCE_ALREADY_EXIST:
      return 'The compagnieAssurance already exists';

    case ErrorMessage.NO_COMPAGNIE_ASSURANCE_FOUND:
      return 'compagnieAssurance is not found';
    default:
      return 'undefind error';
  }
}

export function matchEndorsementError(err: any, errorMessage: string) {
  return err.endorsements.find((e: any) => {
    const matched = e.message === errorMessage;
    if (!matched) {
      console.log(
        `${e.message} is not catched please add a matchEndorsementError on it`,
      );
    }
    return matched;
  });
}

// tslint:disable-next-line:no-any
export function createErrorFromEndorsementError2(
  err: any,
  errorMessage: string,
) {
  // tslint:disable-next-line:no-any
  const extractedError = err.endorsements.find((e: any) => {
    const matched = e.message === errorMessage;
    if (!matched) {
      console.log(
        `${e.message} is not catched please add a matchEndorsementError on it`,
      );
    }
    return matched;
  });
  return new Error(extractedError.message);
}

export function createErrorFromEndorsementError(err: any) {
  if (typeof err.endorsements === 'undefined') {
    return err;
  }
  const extractedError = err.endorsements.find((e: any) => {
    return e.message;
  });
  return new Error(extractedError.message);
}
